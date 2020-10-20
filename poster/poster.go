package poster

import (
	"io/ioutil"
	"net/http"
	"sync"
)

//http header
type PosterHead struct {
	header map[string]string
}

type HtmlCallback func(content string)

type Poster interface {
	Call(url string, head *PosterHead)
}

type HttpPoster struct {
	callbacks   []HtmlCallback
	queryString string
	lock *sync.RWMutex
}


func NewHttpPoster() *HttpPoster {
	return &HttpPoster{callbacks: make([]HtmlCallback, 0)}
}

func (h *HttpPoster) Call(url string, head *PosterHead) {
	resp, err := http.Get(url)
	if err != nil {
		//fmt.Errorf("error", err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		//fmt.Errorf("read body error", err)
	}

	if len(h.callbacks) > 0 {
		for _, cc := range h.callbacks {
			cc(string(body))
		}
	}
}

func (h *HttpPoster) OnHtml(query string, callback HtmlCallback) {
	h.queryString = query
	h.callbacks = append(h.callbacks, callback)
}
