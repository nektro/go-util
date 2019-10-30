package mbpp

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"sync"
	"time"

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
	taskSize  = int64(0)
	barStyle  = "[=>-]<+"
)

var (
	blankWg = new(sync.WaitGroup)
)

func BlankWaitGroup() *sync.WaitGroup {
	blankWg.Add(1)
	return blankWg
}

func Init(concurrency int) {
	doneWg = new(sync.WaitGroup)
	progress = mpb.New(mpb.WithWidth(64), mpb.WithWaitGroup(doneWg))
	guard = semaphore.NewWeighted(int64(concurrency))
}

// SetBarStyle sets the mpb style
//  '1th rune' stands for left boundary rune
// '2th rune' stands for fill rune
//  '3th rune' stands for tip rune
// '4th rune' stands for empty rune
//  '5th rune' stands for right boundary rune
// Default style is `[=>-]`
func SetBarStyle(bstyle string) {
	barStyle = bstyle
}

func CreateJob(name string, f func(*BarProxy, *sync.WaitGroup)) {
	guard.Acquire(ctx, 1)
	func() {
		defer guard.Release(1)

		bar := createBar(name)
		wg := new(sync.WaitGroup)
		f(bar, wg)
		bar.incRaw(1)
		wg.Wait()
		bar.Wait()
	}()
}

func createBar(name string) *BarProxy {
	taskIndex++
	task := fmt.Sprintf("Task #%d", taskIndex)

	b := progress.AddBar(1,
		mpb.BarStyle(barStyle),
		mpb.BarRemoveOnComplete(),
		mpb.PrependDecorators(
			decor.Name(task, decor.WCSyncSpaceR),
			decor.Name(": ", decor.WC{W: 2}),
			decor.Name(name, decor.WCSyncSpaceR),
			decor.Name(": ", decor.WC{W: 2}),
			decor.CountersNoUnit("%d / %d", decor.WCSyncWidth),
		),
		mpb.AppendDecorators(
			decor.OnComplete(decor.Percentage(decor.WCSyncSpace), ""),
			decor.Name(": ", decor.WC{W: 2}),
			decor.OnComplete(decor.EwmaETA(decor.ET_STYLE_MMSS, 0, decor.WCSyncWidth), ""),
		),
	)

	return &BarProxy{1, b, time.Now(), new(sync.WaitGroup)}
}

func Wait() {
	progress.Wait()
}

func GetTaskCount() int {
	return taskIndex
}

func GetTaskDownloadSize() int64 {
	return taskSize
}

func CreateDownloadJob(urlS string, pathS string, mbar *BarProxy) {
	if util.DoesFileExist(pathS) {
		return
	}

	req, err := http.NewRequest(http.MethodGet, urlS, nil)
	if err != nil {
		return
	}
	req.Header.Add("user-agent", "github.com/nektro/go-util/mbpp")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return
	}
	if res.StatusCode != 200 {
		return
	}

	dst, err := os.Create(pathS)
	if err != nil {
		return
	}

	CreateTransferJob(urlS, res.Body, dst, res.ContentLength, mbar)
	res.Body.Close()
	dst.Close()
}

func CreateTransferJob(name string, from io.Reader, to io.Writer, max int64, bar *BarProxy) {
	CreateJob(name, func(b *BarProxy, _ *sync.WaitGroup) {
		if bar != nil {
			defer bar.Increment(1)
		}

		if from == nil || to == nil {
			return
		}
		b.addRaw(max)
		src := &passThru{b, from}
		io.Copy(to, src)
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
	taskSize += int64(n)
	pt.bar.incRaw(n)
	return n, err
}
