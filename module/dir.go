package module

import (
	"fmt"
	"github.com/code-scan/Goal/Ghttp"
	"log"
	"strings"
)

type Scan struct {
	http Ghttp.Http
	Mode string
	Ext  string
}

func NewScan(mode string, ext string) *Scan {
	s := Scan{}
	s.Mode = mode
	s.Ext = ext
	s.http = Ghttp.Http{}
	return &s
}
func (s *Scan) Check(uri string) {
	uri = strings.ReplaceAll(uri, "%EXT%", s.Ext)
	s.http.New(strings.ToUpper(s.Mode), uri)
	s.http.SetUserAgent("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_5) AppleWebKit/537.123 (KHTML, like Gecko) Chrome/80.0.11.4514 Safari/537.123")
	if Proxy != "" {
		s.http.SetProxy(Proxy)
	}
	s.http.Execute()
	if code := s.http.StatusCode(); code != 404 {
		line := fmt.Sprintf("[%d] - [%s]", code, uri)
		log.Println(line)
	}
}
