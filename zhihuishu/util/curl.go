package util

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"zhs/zhihuishu/errorCode"
)

func Curl(method, api, payload string, headers map[string]string) ([]byte, error) {
	client := http.Client{}
	request, newReqErr := http.NewRequest(method, api, strings.NewReader(payload))
	if newReqErr != nil {
		return nil, fmt.Errorf("getCourseInfo -> newReqErr: %s, 错误代码：%d ", newReqErr, errorCode.NewReqError)
	}
	// 设置请求头
	for key, val := range headers {
		request.Header.Set(key, val)
	}
	// 发送请求
	doReq, doReqErr := client.Do(request)
	client.CloseIdleConnections()
	if doReqErr != nil {
		return nil, fmt.Errorf("getCourseInfo -> doReqErr: %s, 错误代码：%d ", doReqErr, errorCode.DoReqError)
	}
	defer func() { _ = doReq.Body.Close() }()
	// 读取返回信息
	reqBody, readBodyErr := ioutil.ReadAll(doReq.Body)
	if readBodyErr != nil {
		return nil, fmt.Errorf("getCourseInfo -> readBodyErr:  %s, 错误代码：%d ", errorCode.IOReadAllError)
	}
	return reqBody, nil
}