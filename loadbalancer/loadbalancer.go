package loadbalancer

import (
	"net/http"
	"net/http/httputil"
	"net/url"
)

type ServerPool struct {
	serverList [] *Server
	indexChannel chan bool
	serverIndex int
}

type Server struct {
	URL *url.URL
	ReverseProxy  *httputil.ReverseProxy
}

func NewServerPool() *ServerPool {
	return &ServerPool{
		indexChannel: make(chan bool, 1),
		serverIndex: 0,
	}
}

func (sp *ServerPool) AddServer(rawURL string) {
	u, _ := url.Parse(rawURL)
	sp.serverList = append(sp.serverList, &Server{
		URL: u,
		ReverseProxy: httputil.NewSingleHostReverseProxy(u),
	})
}

func (sp *ServerPool) GetNextIndex() int {
	temp := -1
	sp.indexChannel <- true
	sp.serverIndex = (sp.serverIndex + 1) % len(sp.serverList)
	temp = sp.serverIndex
	<- sp.indexChannel
	return temp
}

func (sp *ServerPool) GetNextPeer() *Server {
	next := sp.GetNextIndex()
	return sp.serverList[next]
}

func (sp *ServerPool) LoadBalance(w http.ResponseWriter, r *http.Request) {
	peer := sp.GetNextPeer()
	peer.ReverseProxy.ServeHTTP(w, r)
}