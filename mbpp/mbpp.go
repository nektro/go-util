package mbpp

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"sync"

	"github.com/nektro/go-util/util"

	"github.com/vbauerster/mpb"
	"github.com/vbauerster/mpb/decor"
	"golang.org/x/sync/semaphore"
)

var (
	doneWg    *sync.WaitGroup
	progress  *mpb.Progress
	guard     *semaphore.Weighted
	ctx       = context.TODO()
	taskIndex = 0
)

func Init(concurrency int) {
	doneWg = new(sync.WaitGroup)
	progress = mpb.New(mpb.WithWidth(64), mpb.WithWaitGroup(doneWg))
	guard = semaphore.NewWeighted(int64(concurrency))
}

func CreateJob(name string, f func(*BarProxy, *sync.WaitGroup)) {
	guard.Acquire(ctx, 1)
	go func() {
		defer guard.Release(1)

		bar := createBar(name)
		wg := new(sync.WaitGroup)
		f(bar, wg)
		bar.Increment(1)
	}()
}

func createBar(name string) *BarProxy {
	taskIndex++
	task := fmt.Sprintf("Task #%d:", taskIndex)

	b := progress.AddBar(1,
		mpb.BarRemoveOnComplete(),
		mpb.PrependDecorators(
			decor.Name(task, decor.WC{W: len(task) + 1, C: decor.DidentRight}),
			decor.Name(name, decor.WCSyncSpaceR),
			decor.Name(": ", decor.WC{W: 2}),
			decor.CountersNoUnit("%d / %d", decor.WCSyncWidth),
		),
		mpb.AppendDecorators(
			decor.OnComplete(decor.Percentage(decor.WC{W: 5}), ""),
		),
	)

	return &BarProxy{1, b}
}

func Wait() {
	progress.Wait()
}

func GetTaskCount() int {
	return taskIndex
}

func CreateDownloadJob(urlS string, pathS string, wg *sync.WaitGroup) {
	CreateJob(urlS, func(bar *BarProxy, swg *sync.WaitGroup) {
		wg.Add(1)
		defer wg.Done()
		bar.AddToTotal(1)
		defer bar.Increment(1)

		if util.DoesFileExist(pathS) {
			return
		}

		res, err := http.Get(urlS)
		if err != nil {
			return
		}
		if res.StatusCode != 200 {
			return
		}
		bar.AddToTotal(int(res.ContentLength))

		src := &passThru{bar, res.Body}
		dst, err := os.Create(pathS)
		if err != nil {
			return
		}

		io.Copy(dst, src)
		res.Body.Close()
		dst.Close()
	})
}

//
//

type passThru struct {
	bar    *BarProxy
	reader io.Reader
}

func (pt *passThru) Read(p []byte) (int, error) {
	n, err := pt.reader.Read(p)
	pt.bar.Increment(n)
	return n, err
}
