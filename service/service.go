package service

import (
	http2 "ceph/http"
	"github.com/prometheus/alertmanager/template"
	"k8s.io/klog"
)

type Service struct {
	url string
}

func NewService(url string) (result *Service) {
	return &Service{url: url}
}
func (s *Service) Send(prosend template.Data)error{
	for _, alert := range prosend.Alerts {
		delKeys:=[]string{"job","severity","cephid","pool_id"}
		labels:=alert.Labels.Remove(delKeys)
		klog.Infof("Alert: status=%s,Labels=%v,Annotations=%v", alert.Status, labels, alert.Annotations)
		alert.Labels=labels
		if err := s.PrometheusSend(alert); err != nil {
			return err
		}
	}
	return nil
}
func (s *Service) PrometheusSend(proSend template.Alert) error {

	msg := http2.SendMsg{
		Msgtype: "markdown",
	}
	url := s.url
	data:=http2.RobotMsgModel(proSend)
	msg.Markdown = &http2.MsgContent{Content: data}
	err:=http2.DoPost(url,msg)
	if err!=nil{
		klog.Errorf("Webhook:send url false: %v", err.Error())
		return err
	}
	return nil
}

