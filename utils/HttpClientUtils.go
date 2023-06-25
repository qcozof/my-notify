package utils

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type ResponseWrapper struct {
	StatusCode int
	Body       string
	Header     http.Header
}

func Get(url string, token string, timeout int) ResponseWrapper {
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("x-token", token)
	if err != nil {
		return createRequestError(err)
	}

	return httprequest(req, timeout)
}

func PostParams(url string, params string, timeout int) ResponseWrapper {
	buf := bytes.NewBufferString(params)
	req, err := http.NewRequest("POST", url, buf)
	if err != nil {
		return createRequestError(err)
	}
	req.Header.Set("Content-type", "application/x-www-form-urlencoded")

	return httprequest(req, timeout)
}

func PostJson(url string, body string, timeout int) ResponseWrapper {
	buf := bytes.NewBufferString(body)
	req, err := http.NewRequest("POST", url, buf)
	if err != nil {
		return createRequestError(err)
	}
	req.Header.Set("Content-type", "application/json")

	return httprequest(req, timeout)
}

func httprequest(req *http.Request, timeout int) ResponseWrapper {
	wrapper := ResponseWrapper{StatusCode: 0, Body: "", Header: make(http.Header)}
	client := &http.Client{}
	if timeout > 0 {
		client.Timeout = time.Duration(timeout) * time.Second
	}
	setRequestHeader(req)
	resp, err := client.Do(req)
	if err != nil {
		wrapper.Body = fmt.Sprintf("执行HTTP请求错误-%s", err.Error())
		return wrapper
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		wrapper.Body = fmt.Sprintf("读取HTTP请求返回值失败-%s", err.Error())
		return wrapper
	}
	wrapper.StatusCode = resp.StatusCode
	wrapper.Body = string(body)
	wrapper.Header = resp.Header

	return wrapper
}

func setRequestHeader(req *http.Request) {
	req.Header.Set("User-Agent", "golang/gocron")
}

func createRequestError(err error) ResponseWrapper {
	errorMessage := fmt.Sprintf("创建HTTP请求错误-%s", err.Error())
	return ResponseWrapper{0, errorMessage, make(http.Header)}
}

//检测是否返回200
func CheckUrlHttpCodeOk(url, method string) (bool, error) {

	req, err := http.NewRequest(strings.ToUpper(method), url, nil)
	if err != nil {
		return false, err
	}
	resp := httprequest(req, 30)
	ok := resp.StatusCode == 200
	if !ok {
		//fmt.Println(resp.Body)
	}
	return ok, errors.New(fmt.Sprintf("http状态码：%d", resp.StatusCode))
}

func GetValFromQuery(queryStr, key string) (string, error) {

	u, err := url.Parse(queryStr)
	if err != nil {
		return "", errors.New("参数不合法。")
	}

	m, _ := url.ParseQuery(u.RawQuery)

	if len(m[key]) > 0 {
		return m[key][0], nil
	}
	return "",nil
}
