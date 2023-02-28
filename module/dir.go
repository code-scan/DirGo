package module

import (
	"fmt"
	"log"
	"strings"

	"github.com/code-scan/Goal/Ghttp"
)

type Scan struct {
	http    Ghttp.Http
	Mode    string
	Ext     string
	Keyword string
	Codes   map[int]bool
}

func NewScan(mode string, ext, keyword string, codes map[int]bool) *Scan {
	s := Scan{}
	s.Mode = strings.ToUpper(mode)
	s.Ext = ext
	s.Keyword = keyword
	s.Codes = codes
	s.http = Ghttp.Http{}
	return &s
}
func (s *Scan) Check(uri string) {
	uri = strings.ReplaceAll(uri, "%EXT%", s.Ext)
	s.http.New(s.Mode, uri)
	s.http.SetUserAgent("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_5) AppleWebKit/537.123 (KHTML, like Gecko) Chrome/80.0.11.4514 Safari/537.123")
	if Proxy != "" {
		s.http.SetProxy(Proxy)
	}
	s.http.Execute()
	var text = ""
	if s.Mode != "HEAD" {
		text, _ = s.http.Text()
	}
	defer s.http.Close()
	if code := s.http.StatusCode(); code > 0 && (s.Mode != "HEAD" && !strings.Contains(text, s.Keyword)) {
		if ok, _ := s.Codes[code]; !ok {
			return
		}
		line := fmt.Sprintf("[%d] - [ %s ]", code, uri)
		log.Println(line)
	}
}
