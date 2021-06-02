package service

import (
	http2 "ceph/http"
	"github.com/prometheus/alertmanager/template"
	"k8s.io/klog"
	"strconv"
)

type Service struct {
	url string
}

func NewService(url string) (result *Service) {
	return &Service{url: url}
}
func (s *Service) Send(prosend template.Data) error {
	// o`nly use prosend.Alert[0]
	if err:=s.RobotGroupMsg(len(prosend.Alerts),prosend.Alerts[0],prosend.Alerts[0].Labels[http2.LabelsKeySeverity]);err!=nil{
		klog.Errorf("Send-RobotGroupMsg is error", err.Error())
		return err
	}
	return nil
}

func (s *Service) RobotGroupMsg(len int, alerts template.Alert,addDes string) error {
	alerts.Labels=alerts.Labels.Remove(http2.DelKeys)
	content, err := http2.RobotMsgModelOne(alerts,addDes)
	if err != nil {
		klog.Errorf("RobotGroupMsg-RobotMsgModelOne is error", err.Error())
		return err
	}
	if err:=s.PrometheusSend(content,len);err!=nil{
		klog.Errorf("RobotGroupMsg-PrometheusSend is error", err.Error())
		return err
	}
	return nil
}

func (s *Service) PrometheusSend(msgContent string,len int) error {
	msg := http2.SendMsg{
		Msgtype: "markdown",
	}
	url := s.url
	msgContent+="\n"+"还有"+strconv.Itoa(len-1)+"条数没查看"
	msg.Markdown = &http2.MsgContent{Content: msgContent}
	err := http2.DoPost(url, msg)
	if err != nil {
		klog.Errorf("Webhook:send url false: %v", err.Error())
		return err
	}
	return nil
}
