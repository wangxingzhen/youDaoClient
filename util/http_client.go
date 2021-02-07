package util

import (
	"fmt"
	json "github.com/json-iterator/go"
	"github.com/valyala/fasthttp"
)

// get请求
func GetJson(url string, header map[string]string, response interface{}) (int, error) {
	code, body, err := getQuery(url, header)
	if err != nil {
		return 0, err
	}
	return code, json.Unmarshal(body, response)
}

// get请求
func getQuery(url string, header map[string]string) (int, []byte, error) {
	req := fasthttp.AcquireRequest()
	req.Header.SetMethod("GET")
	if len(header) > 0 {
		for key, val := range header {
			req.Header.Add(key, val)
		}
	}
	req.SetRequestURI(url)
	resp := fasthttp.AcquireResponse()
	if err := fasthttp.Do(req, resp); err != nil {
		return 0, nil, err
	}
	var respCode, respBodyTmp = resp.StatusCode(), resp.Body()
	respBody := make([]byte, len(respBodyTmp))
	copy(respBody, respBodyTmp)
	fasthttp.ReleaseResponse(resp)
	fasthttp.ReleaseRequest(req)
	return respCode, respBody, nil
}

// json请求
func PostJson(url string, params interface{}, header map[string]string, response interface{}) (int, error) {
	raw, err := json.Marshal(params)
	if err != nil {
		return 0, err
	}
	code, body, err := postQuery(url, raw, header)
	if err != nil {
		return 0, err
	}
	return code, json.Unmarshal(body, response)
}

// post 请求
func postQuery(url string, params []byte, header map[string]string) (int, []byte, error) {
	req := fasthttp.AcquireRequest()
	req.Header.SetContentType("application/json; charset=UTF-8")
	req.Header.SetMethod("POST")
	if len(header) > 0 {
		for key, val := range header {
			req.Header.Add(key, val)
		}
	}
	req.SetRequestURI(url)
	requestBody := params
	req.SetBody(requestBody)
	resp := fasthttp.AcquireResponse()
	if err := fasthttp.Do(req, resp); err != nil {
		return 0, nil, err
	}
	var respCode, respBodyTmp = resp.StatusCode(), resp.Body()
	respBody := make([]byte, len(respBodyTmp))
	copy(respBody, respBodyTmp)
	fasthttp.ReleaseResponse(resp)
	fasthttp.ReleaseRequest(req)
	return respCode, respBody, nil
}

// from-data请求
func PostForm(url string, params map[string]string, header map[string]string, response interface{}) (int, error) {
	args := &fasthttp.Args{}
	for i, i2 := range params {
		args.Add(i, i2)
	}
	code, body, err := postFormQuery(url, args, header)
	fmt.Println(string(body))
	if err != nil {
		return 0, err
	}
	return code, json.Unmarshal(body, response)
}

// postForm 请求
func postFormQuery(url string, params *fasthttp.Args, header map[string]string) (int, []byte, error) {
	req := &fasthttp.Request{}
	req.Header.SetMethod("POST")
	if len(header) > 0 {
		for key, val := range header {
			req.Header.Add(key, val)
		}
	}
	req.Header.SetContentType("application/x-www-form-urlencoded; charset=UTF-8")
	status, resp, err := fasthttp.Post(nil, url, params)
	if err != nil {
		return 0, nil, err
	}
	var respCode, respBodyTmp = status, resp
	respBody := make([]byte, len(respBodyTmp))
	copy(respBody, respBodyTmp)
	return respCode, respBody, nil
}
