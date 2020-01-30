package util

import (
	"bufio"
	"errors"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"os/signal"
	"runtime/debug"
	"strings"
	"syscall"
	"time"

	"github.com/nektro/go-util/ansi/style"

	. "github.com/nektro/go-util/alias"
)

func Log(message ...interface{}) {
	fmt.Print(GetIsoDateTime() + ": ")
	fmt.Println(message...)
}

func LogError(message ...interface{}) {
	fmt.Print(style.FgRed)
	Log(message...)
	fmt.Print(style.ResetAll)
}

func LogWarn(message ...interface{}) {
	fmt.Print(style.FgYellow)
	Log(message...)
	fmt.Print(style.ResetAll)
}

func GetIsoDateTime() string {
	return time.Now().UTC().String()[0:19]
}

func PrintTable(data [][]string, dividers bool) {
	rows := countRows(data)
	cols := countColumns(data)
	widths := countMaxColumWidths(data, rows, cols)

	fmt.Print("┌")
	for i := 0; i < len(widths); i++ {
		fmt.Print(strings.Repeat("─", widths[i]))
		if i < len(widths)-1 {
			fmt.Print("┬")
		} else {
			fmt.Print("┐")
		}
	}
	fmt.Println()

	for i := 0; i < rows; i++ {
		fmt.Print("│")
		for j := 0; j < cols; j++ {
			if j < len(data[i]) {
				fmt.Print(data[i][j])
				fmt.Print(strings.Repeat(" ", widths[j]-len(data[i][j])))
			} else {
				fmt.Print(strings.Repeat(" ", widths[j]))
			}
			fmt.Print("│")
		}
		fmt.Println()

		if dividers {
			if i < rows-1 {
				fmt.Print("├")
				for j := 0; j < cols; j++ {
					fmt.Print(strings.Repeat("─", widths[j]))
					if j < cols-1 {
						fmt.Print("┼")
					} else {
						fmt.Print("┤")
					}
				}
				fmt.Println()
			}
		}
	}

	fmt.Print("└")
	for i := 0; i < cols; i++ {
		fmt.Print(strings.Repeat("─", widths[i]))
		if i < cols-1 {
			fmt.Print("┴")
		} else {
			fmt.Print("┘")
		}
	}
	fmt.Println()
}

func countRows(data [][]string) int {
	return len(data)
}

func countColumns(data [][]string) int {
	res := 0
	for _, item := range data {
		l := len(item)
		if l > res {
			res = l
		}
	}
	return res
}

func countMaxColumWidths(data [][]string, rowCount int, colCount int) []int {
	res := make([]int, colCount)
	for i := 0; i < rowCount; i++ {
		for j := 0; j < colCount; j++ {
			if j == len(data[i]) {
				break
			}
			l := len(data[i][j])
			if l > res[j] {
				res[j] = l
			}
		}
	}
	return res
}

func DieOnError(err error, args ...string) {
	if err != nil {
		LogError(err.Error())
		for _, item := range args {
			LogError(item)
		}
		os.Exit(1)
	}
}

func Assert(condition bool, errorMessage string) error {
	if condition {
		return nil
	}
	return errors.New(errorMessage)
}

func DoesFileExist(file string) bool {
	_, err := os.Stat(file)
	return !os.IsNotExist(err)
}

func DoesDirectoryExist(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	if !s.IsDir() {
		return false
	}
	return true
}

func ReadFile(path string) []byte {
	reader, _ := os.Open(path)
	bytes, _ := ioutil.ReadAll(reader)
	return bytes
}

func ReadFileLines(path string, send func(string)) {
	reader, _ := os.Open(path)
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		send(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		LogError(err)
	}
}

func CheckErr(err error, args ...string) {
	if err != nil {
		LogError(F("%q: %s", err, args))
		debug.PrintStack()
	}
}

func Contains(haystack []string, needle string) bool {
	for _, item := range haystack {
		if needle == item {
			return true
		}
	}
	return false
}

func FullHost(r *http.Request) string {
	urlS := "http"
	if len(r.Header.Get("X-TLS-Enabled")) > 0 {
		urlS += "s"
	}
	return urlS + "://" + r.Host
}

func DoHttpRequest(req *http.Request) []byte {
	resp, _ := http.DefaultClient.Do(req)
	body, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	return body
}

func IsPortAvailable(port int) bool {
	ln, err := net.Listen("tcp", F(":%d", port))
	if err != nil {
		return false
	}
	ln.Close()
	return true
}

func Btoi(b bool) int {
	if b {
		return 1
	}
	return 0
}

func ReduceNumber(input int64, unit int64, base string, prefixes string) string {
	if input < unit {
		return F("%d "+base, input)
	}
	div, exp := int64(unit), 0
	for n := input / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return F("%.1f %ci", float64(input)/float64(div), prefixes[exp]) + base
}

func ByteCountIEC(b int64) string {
	return ReduceNumber(b, 1024, "B", "KMGTPEZY")
}

func RunOnClose(f func()) {
	gracefulStop := make(chan os.Signal)
	signal.Notify(gracefulStop, syscall.SIGTERM)
	signal.Notify(gracefulStop, syscall.SIGINT)

	go func() {
		sig := <-gracefulStop
		fmt.Println()
		Log(F("Caught signal '%+v'", sig))
		f()
		os.Exit(0)
	}()
}

func TrimLen(s string, l int) string {
	if len(s) <= l {
		return s
	}
	return s[:l]
}

func FirstNonZero(x ...int) int {
	for _, item := range x {
		if item != 0 {
			return item
		}
	}
	return 0
}

func FirstNonEmptyS(values ...string) string {
	for _, item := range values {
		if len(item) > 0 {
			return item
		}
	}
	return ""
}

func DoHttpFetch(req *http.Request) ([]byte, error) {
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
