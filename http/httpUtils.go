package http2

import (
	"bytes"
	"encoding/json"
	"k8s.io/klog"
	"net/http"
)

func DoPost(url string, msg interface{}) error {
	jsonBytes, err := json.Marshal(msg)
	if err != nil {
		klog.Errorf("sendToWechatWork: json marshal error: %v", err.Error())
		return err
	}

	reader := bytes.NewReader(jsonBytes)
	request, err := http.NewRequest("POST", url, reader)
	if err != nil {
		klog.Errorf("sendToWechatWork: http new request error: %v", err.Error())
		return err
	}

	request.Header.Set("Content-Type", "application/json")
	client := http.Client{}
	_, err = client.Do(request)
	if err != nil {
		klog.Errorf("sendToWechatWork: http post request error: %v", err.Error())
		return err
	}

	return nil
}
