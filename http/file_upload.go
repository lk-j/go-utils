package http

import (
	"bytes"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
)
 /*
  *  模拟客户端上传单文件
  * filename 文件名
  * targetUrl 目标地址
  * 文件参数
  */
func PostFile(filename,targetUrl,query string) (string,error) {
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)

	//关键的一步操作
	fileWriter, err := bodyWriter.CreateFormFile(query, filename)
	if err != nil {
		return "",err
	}

	//打开文件句柄操作
	fh, err := os.Open(filename)
	if err != nil {
		return "",err
	}
	defer fh.Close()

	//iocopy
	_, err = io.Copy(fileWriter, fh)
	if err != nil {
		return "",err
	}

	contentType := bodyWriter.FormDataContentType()
	_ = bodyWriter.Close()

	resp, err := http.Post(targetUrl, contentType, bodyBuf)
	if err != nil {
		return "",err
	}
	defer resp.Body.Close()
	resp_body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "",err
	}
	return string(resp_body),nil
}
