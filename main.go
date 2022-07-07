package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func modifyResponse() func(*http.Response) error {
	return func(resp *http.Response) error {
		switch resp.Request.URL.Path {
		case "/block_count":
		case "/mine":
		case "/tx_json":
		default:
			b, _ := httputil.DumpResponse(resp, true)
			log.Println("<- ", resp.Request.URL.Path)
			println(string(b))
			log.Println("----")
		}
		return nil
	}
}

func main() {
	remote, err := url.Parse("http://65.108.201.140:3000")
	if err != nil {
		panic(err)
	}

	handler := func(p *httputil.ReverseProxy) func(http.ResponseWriter, *http.Request) {
		return func(w http.ResponseWriter, r *http.Request) {
			switch r.URL.Path {
			case "/block_count":
			case "/mine":
			case "/tx_json":
			default:
				b, _ := httputil.DumpRequest(r, true)
				log.Println("-> ", r.URL.Path)
				println(string(b))
				log.Println("----")
			}

			r.Host = remote.Host
			p.ServeHTTP(w, r)
		}
	}

	proxy := httputil.NewSingleHostReverseProxy(remote)
	proxy.ModifyResponse = modifyResponse()
	http.HandleFunc("/", handler(proxy))
	err = http.ListenAndServe(":3000", nil)
	if err != nil {
		panic(err)
	}
}
