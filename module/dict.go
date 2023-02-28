package module

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"sync"
	"time"
)

var Proxy string
var TaskQueue = make(chan string)
var Wg = &sync.WaitGroup{}

func ReadListToArray(f string, a *[]string) {
	if data, err := ioutil.ReadFile(f); err != nil {
		log.Println("file path is error ", err)
	} else {
		d := strings.Split(string(data), "\n")
		for _, p := range d {
			if p == "" {
				continue
			}
			p = strings.TrimSpace(p)
			*a = append(*a, p)
		}
	}
}

func AddTask(targetList, dirList []string) {
	for _, t := range targetList {
		if t[0:4] != "http" {
			t = fmt.Sprintf("http://%s", t)
		}
		ulen := len(t)
		for _, d := range dirList {
			var uri string
			if len(d) > 0 && d[0:1] != "/" && t[ulen-1:] != "/" {
				uri = fmt.Sprintf("%s/%s", t, d)
			} else {
				uri = fmt.Sprintf("%s%s", t, d)
			}
			TaskQueue <- uri
		}
	}
}
func RunTask(mode, ext, keyword string, codes map[int]bool) {
	for {
		scan := NewScan(mode, ext, keyword, codes)
		select {
		case u := <-TaskQueue:
			scan.Check(u)
		case <-time.After(5 * time.Second):
			Wg.Done()
		}
	}
}
