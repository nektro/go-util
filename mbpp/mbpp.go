package mbpp

import (
	"context"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/nektro/go-util/ansi/style"
	"github.com/nektro/go-util/util"
	"github.com/vbauerster/mpb"
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

func CreateJob(name string, f func(*BarProxy)) {
	tryToStartJob(0, name, f)
}

func tryToStartJob(typ int, name string, f func(*BarProxy)) {
	guard.Acquire(ctx, 1)
	func() {
		defer guard.Release(1)

		bar := createBar(name, typ)
		f(bar)
		bar.Wait()
		bar.incRaw(1)
	}()
}

func createBar(name string, typ int) *BarProxy {
	taskIndex++
	task := style.FgGreen + "Task " + strconv.Itoa(taskIndex) + style.ResetFgColor

	switch typ {
	case 0:
		return &BarProxy{1, barA(task, name), time.Now(), new(sync.WaitGroup)}
	case 1:
		return &BarProxy{1, barB(task, name), time.Now(), new(sync.WaitGroup)}
	}
	return nil
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

func GetCompletionMessage() string {
	return "Complete after " + strconv.Itoa(GetTaskCount()) + " tasks and " + util.ByteCountIEC(GetTaskDownloadSize()) + " downloaded."
}

func updateBar(bar *BarProxy) {
	if bar != nil {
		bar.Increment(1)
	}
}

func CreateDownloadJob(urlS string, pathS string, mbar *BarProxy) {
	tryToStartJob(1, urlS, func(bar *BarProxy) {
		defer updateBar(mbar)

		if util.DoesFileExist(pathS) {
			r, _ := http.Head(urlS)
			f, _ := os.Open(pathS)
			s, _ := f.Stat()
			if r != nil && s != nil {
				if s.Size() == r.ContentLength && r.ContentLength != 0 {
					return
				}
			}
		}

		res, err := httpReqWithRetry(urlS)
		if err != nil {
			return
		}
		defer res.Body.Close()
		if res.StatusCode != 200 {
			return
		}

		dst, err := os.Create(pathS)
		if err != nil {
			return
		}
		defer dst.Close()

		DoBarTransfer(res.Body, dst, res.ContentLength, bar)
	})
}

func DoBarTransfer(from io.Reader, to io.Writer, max int64, bar *BarProxy) {
	if from == nil || to == nil {
		return
	}
	if max > 0 {
		bar.addRaw(max)
	}
	src := &passThru{bar, from, max > 0}
	io.Copy(to, src)
}

func CreateHeadlessJob(name string, max int64, bar *BarProxy) *HeadlessBar {
	r, w := io.Pipe()
	go CreateJob(name, func(b *BarProxy) {
		defer updateBar(bar)
		DoBarTransfer(r, ioutil.Discard, max, b)
	})
	return &HeadlessBar{w, max, 0}
}

//
//

type passThru struct {
	bar    *BarProxy
	reader io.Reader
	determ bool
}

func (pt *passThru) Read(p []byte) (int, error) {
	n, err := pt.reader.Read(p)
	taskSize += int64(n)
	if !pt.determ {
		pt.bar.addRaw(int64(n))
	}
	pt.bar.incRaw(n)
	return n, err
}
