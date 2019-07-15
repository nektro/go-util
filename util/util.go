package util

import (
	"bufio"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"runtime/debug"
	"strings"
	"time"

	"github.com/fatih/color"
	"github.com/valyala/fasthttp"

	. "github.com/nektro/go-util/alias"
)

var (
	fasthttpHandlers = map[string]func(ctx *fasthttp.RequestCtx){}
)

func Log(message ...interface{}) {
	fmt.Print("[" + GetIsoDateTime() + "] ")
	fmt.Println(message...)
}

func LogError(message ...interface{}) {
	color.Red("["+GetIsoDateTime()+"] [error] %s", message...)
}

func Logf(format string, args ...interface{}) {
	fmt.Println("[" + GetIsoDateTime() + "] " + F(format, args...))
}

func GetIsoDateTime() string {
	return time.Now().UTC().String()[0:19]
}

func FasthttpAddHandler(path string, handle func(ctx *fasthttp.RequestCtx)) {
	fasthttpHandlers[path] = handle
}

func FasthttpHandle(path string, ctx *fasthttp.RequestCtx) bool {
	for k, v := range fasthttpHandlers {
		if k == path {
			v(ctx)
			return true
		}
	}
	return false
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
		LogError(err)
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
		os.Exit(2)
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
	urL := "http"
	if r.TLS != nil {
		urL += "s"
	}
	return urL + "://" + r.Host
}

func DoHttpRequest(req *http.Request) []byte {
	resp, _ := http.DefaultClient.Do(req)
	body, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	return body
}
