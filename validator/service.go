package validator

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/sirupsen/logrus"

	"code.as/core/socks"
)

type Service interface {
	// Check and return onlines proxies
	Check(proxies []string) ProxiesResult
}

func NewService(timeout time.Duration) Service {
	return &service{
		TestURL: "https://api.my-ip.io/ip",
		Timeout: timeout,
	}
}

type service struct {
	// TestURL without protocol
	// "api.my-ip.io/ip"
	TestURL string

	Timeout time.Duration
}

// Proxy format examples: "socks4://200.01.01.01:4153", "http://100.100.100.1:80"
func (s service) Check(proxies []string) ProxiesResult {
	results := make(chan ProxyResult, len(proxies))

	// Verify many proxies at same time
	for _, proxy := range proxies {
		go func(p string) {
			result, _ := s.checkProxy(p)
			results <- result
		}(proxy)
	}

	checkedProxies := []ProxyResult{}
	for range proxies {
		result := <-results
		checkedProxies = append(checkedProxies, result)
	}

	return checkedProxies
}

func (s service) getProxyTransport(proxy string) (*http.Transport, error) {
	proxyURLSplitted := strings.Split(proxy, "://")
	proxyProtocol := proxyURLSplitted[0]
	proxyAddress := proxyURLSplitted[1]

	switch proxyProtocol {
	case "socks4":
		return &http.Transport{
			Dial: socks.DialSocksProxy(socks.SOCKS4, proxyAddress),
		}, nil
	case "socks5":
		return &http.Transport{
			Dial: socks.DialSocksProxy(socks.SOCKS5, proxyAddress),
		}, nil
	case "http", "https":
		proxyURL, err := url.Parse(proxy)
		if err != nil {
			return nil, err
		}

		return &http.Transport{
			Proxy: http.ProxyURL(proxyURL),
		}, nil
	default:
		return nil, fmt.Errorf("%s proxy protocol not supported", proxyProtocol)
	}
}

// checkProxy verify conection with one proxy
func (s service) checkProxy(proxy string) (ProxyResult, error) {
	start := time.Now()

	transport, err := s.getProxyTransport(proxy)
	if err != nil {
		logrus.Errorln("Error on get proxy transport", err)
		return ProxyResult{}, err
	}

	client := resty.New().
		SetTimeout(s.Timeout).
		SetTransport(transport)

	resp, err := client.R().Get(s.TestURL)
	if err != nil {
		logrus.Errorln("Error on get method", err)
	}

	elapsed := time.Since(start)

	return ProxyResult{
		Proxy:    proxy,
		Latency:  elapsed,
		IsOnline: err == nil && resp.StatusCode() == 200,
		Result:   resp.String(),
	}, nil
}
