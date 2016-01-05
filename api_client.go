package godragon

import (
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"
)

/*
BR	BR1	br.api.pvp.net
EUNE	EUN1	eune.api.pvp.net
EUW	EUW1	euw.api.pvp.net
KR	KR	kr.api.pvp.net
LAN	LA1	lan.api.pvp.net
LAS	LA2	las.api.pvp.net
NA	NA1	na.api.pvp.net
OCE	OC1	oce.api.pvp.net
TR	TR1	tr.api.pvp.net
RU	RU	ru.api.pvp.net
PBE	PBE1	pbe.api.pvp.net
*/

// APIClient Riot API client
type APIClient struct {
	endpoint string
	game     string
	region   string
	client   *http.Client
	key      string // API key
	tokens   chan struct{}
}

// NewAPIClient create an initalized APIClient
func NewAPIClient(region, key string) *APIClient {
	a := &APIClient{
		key:    key,
		game:   "lol",
		region: region,
		client: &http.Client{
			Jar:     nil,
			Timeout: time.Second * 5,
			Transport: &http.Transport{
				Proxy: http.ProxyFromEnvironment,
				Dial: (&net.Dialer{
					Timeout:   30 * time.Second,
					KeepAlive: 30 * time.Second,
				}).Dial,
				TLSHandshakeTimeout: 10 * time.Second,
			}},
		tokens: make(chan struct{}, 20),
	}

	return a
}

func (a *APIClient) genURL(version, api string, query url.Values) url.URL {
	u := url.URL{}
	u.Scheme = "https"
	u.Host = strings.Join([]string{a.region, ".api.pvp.net"}, "")
	u.Path = fmt.Sprintf("/api/%s/%s/%s/%s", a.game, a.region, version, api)
	u.RawQuery = query.Encode()
	return u
}

// do execute a request
func (a *APIClient) do(req *http.Request) ([]byte, error) {
	// add api key

	// manipulate request
	query := req.URL.Query()
	query.Add("api_key", a.key)

	//rebuild request

	req.URL.RawQuery = query.Encode()

	rc := make(chan struct {
		data []byte
		err  error
	})

	go func(req http.Request) {
		defer func() {
			<-a.tokens
		}()

		a.tokens <- struct{}{}
		resp, err := a.client.Do(&req)
		if err != nil {
			rc <- struct {
				data []byte
				err  error
			}{nil, err}
		}
		defer resp.Body.Close()

		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			rc <- struct {
				data []byte
				err  error
			}{nil, err}
		}

		rc <- struct {
			data []byte
			err  error
		}{data, err}
	}(*req)

	r := <-rc
	return r.data, r.err
}
