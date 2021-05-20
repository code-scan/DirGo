package main

import (
	"flag"
	"fmt"
	"github.com/code-scan/DirGo/module"
	"log"
	"os"
)

var (
	target      string
	batchTarget string
	dirFile     string
	ext         string
	mode        string
	outFile     string
	maxThread   int
)
func banner(){
	fmt.Println("\n██████╗ ██╗██████╗  ██████╗  ██████╗ \n██╔══██╗██║██╔══██╗██╔════╝ ██╔═══██╗\n██║  ██║██║██████╔╝██║  ███╗██║   ██║\n██║  ██║██║██╔══██╗██║   ██║██║   ██║\n██████╔╝██║██║  ██║╚██████╔╝╚██████╔╝\n╚═════╝ ╚═╝╚═╝  ╚═╝ ╚═════╝  ╚═════╝ \n                                     \n")
	fmt.Println("@Cond0r http://aq.mk\n\n")

}
func main() {
	banner()
	flag.StringVar(&target, "t", "", "target url ,http://baidu.com")
	flag.StringVar(&batchTarget, "f", "", "target file ")
	flag.StringVar(&dirFile, "d", "dict/dirs.txt", "target file ")
	flag.StringVar(&ext, "e", "php ", "ext , php,jsp")
	flag.StringVar(&mode, "m", "HEAD", "GET/HEAD/POST")
	//flag.StringVar(&outFile, "o", "log.txt", "output file")
	flag.StringVar(&module.Proxy, "x", "", "proxy, socks5://user:pass@host:port, http://host:port")

	flag.IntVar(&maxThread, "b", 20, "max thread")
	flag.Parse()

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
		go module.RunTask(mode, ext)
	}
	module.Wg.Wait()
}
