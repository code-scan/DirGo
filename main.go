package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/code-scan/DirGo/module"
	"github.com/code-scan/Goal/Gconvert"
)

var (
	target      string
	batchTarget string
	dirFile     string
	ext         string
	mode        string
	outFile     string
	keyWord     string
	Codes       map[int]bool
	codes       string
	maxThread   int
)

func banner() {
	fmt.Println("\n██████╗ ██╗██████╗  ██████╗  ██████╗ \n██╔══██╗██║██╔══██╗██╔════╝ ██╔═══██╗\n██║  ██║██║██████╔╝██║  ███╗██║   ██║\n██║  ██║██║██╔══██╗██║   ██║██║   ██║\n██████╔╝██║██║  ██║╚██████╔╝╚██████╔╝\n╚═════╝ ╚═╝╚═╝  ╚═╝ ╚═════╝  ╚═════╝ \n                                     \n")
	fmt.Println("@Cond0r http://aq.mk\n\n")

}
func main() {
	banner()
	flag.StringVar(&target, "u", "", "target url ,http://baidu.com")
	flag.StringVar(&keyWord, "k", "", "404 keyword")
	flag.StringVar(&batchTarget, "f", "", "target file ")
	flag.StringVar(&dirFile, "d", "dict/dirs.txt", "dict file ")
	flag.StringVar(&codes, "c", "200,302,301,403", "status code")
	flag.StringVar(&ext, "e", "php ", "ext , php,jsp")
	flag.StringVar(&mode, "m", "HEAD", "GET/HEAD/POST")
	//flag.StringVar(&outFile, "o", "log.txt", "output file")
	flag.StringVar(&module.Proxy, "x", "", "proxy, socks5://user:pass@host:port, http://host:port")

	flag.IntVar(&maxThread, "b", 20, "max thread")
	flag.Parse()
	cs := strings.Split(codes, ",")
	Codes = make(map[int]bool)
	for _, c := range cs {
		Codes[Gconvert.Str2Int(c)] = true
	}

	if target == "" && batchTarget == "" {
		log.Println("pls set target")
		log.Println("usage:")
		log.Println(os.Args[0], " -h")
		return
	}
	var dirList []string
	var targetList []string
	module.ReadListToArray(dirFile, &dirList)
	if batchTarget != "" {
		module.ReadListToArray(batchTarget, &targetList)
	} else {
		targetList = append(targetList, target)
	}
	module.TaskQueue = make(chan string, maxThread*30)
	go module.AddTask(targetList, dirList)
	for i := 0; i < maxThread; i++ {
		module.Wg.Add(1)
		go module.RunTask(mode, ext, keyWord, Codes)
	}
	module.Wg.Wait()
}
