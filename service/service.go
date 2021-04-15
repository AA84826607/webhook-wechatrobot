package service

import (
	"bytes"
	http2 "ceph/http"
	"github.com/prometheus/alertmanager/template"
	"k8s.io/klog"
	gotemplate "text/template"
)

type Service struct {
}

func NewService() (result *Service) {
	return &Service{}
}
func (s *Service) Send(prosend template.Data)error{
	for _, alert := range prosend.Alerts {
		klog.Infof("Alert: status=%s,Labels=%v,Annotations=%v", alert.Status, alert.Labels, alert.Annotations)
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
	url := ""

	var doc bytes.Buffer
	t, err := gotemplate.New("alert").Parse(http2.Templ)
	if err != nil {
		klog.Errorf("Webhook: initial go template error: %v", err.Error())
		return err
	}
	if err := t.Execute(&doc, proSend); err != nil {
		klog.Errorf("Webhook: go template execute error: %v", err.Error())
		return err
	}
	msg.Markdown = &http2.MsgContent{Content: doc.String()}
	err=http2.DoPost(url,msg)
	if err!=nil{
		klog.Errorf("Webhook:send url false: %v", err.Error())
		return err
	}

	return nil
}
