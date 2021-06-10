package service

import (
	http2 "ceph/http"
	"github.com/prometheus/alertmanager/template"
	"k8s.io/klog"
	"strconv"
)

const WechatUrl = "https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key="

type Service struct {
	key string
}

func NewService(key string) (result *Service) {
	return &Service{key: key}
}
func (s *Service) Send(prosend template.Data) error {
	// o`nly use prosend.Alert[0]
	if err := s.RobotGroupMsg(len(prosend.Alerts), prosend.Alerts[0], prosend.Alerts[0].Labels[http2.LabelsKeySeverity]); err != nil {
		klog.Errorf("Send-RobotGroupMsg is error", err.Error())
		return err
	}
	return nil
}
//这是zhuyou的代码 ， 有问题就恢复这里
//func (s *Service) RobotGroupMsg(len int, alerts template.Alert, addDes string) error {
//	alerts.Labels = alerts.Labels.Remove(http2.DelKeys)
//	content, err := http2.RobotMsgModelOne(alerts, addDes)
//	if err != nil {
//		klog.Errorf("RobotGroupMsg-RobotMsgModelOne is error", err.Error())
//		return err
//	}
//	if err := s.PrometheusSend(content, len); err != nil {
//		klog.Errorf("RobotGroupMsg-PrometheusSend is error", err.Error())
//		return err
//	}
//	return nil
//}
//啦啦啦

//这是测试ceph和lotus分组的代码
func (s *Service) RobotGroupMsg(len int, alerts template.Alert, addDes string) error {
	alerts.Labels = alerts.Labels.Remove(http2.DelKeys)
	switch  s.key {
	case   "441670ee-497b-4471-b066-6a2c1ab10c41":
		// 编辑模板
		content, err := http2.RobotMsgModelOne(alerts, addDes)
		if err != nil {
			klog.Errorf("RobotGroupMsg-RobotMsgModelOne is error", err.Error())
			return err
		}
		// 发送给企业微信机器人
		if err := s.PrometheusSend(content, len); err != nil {
			klog.Errorf("RobotGroupMsg-PrometheusSend is error", err.Error())
			return err
		}
	case   "5be934ae-dea8-4085-a6f2-5012e6528dd0":
		// 编辑模板
		content, err := http2.RobotMsgModelTwo(alerts, addDes)
		if err != nil {
			klog.Errorf("RobotGroupMsg-RobotMsgModelOne is error", err.Error())
			return err
		}
		// 发送给企业微信机器人
		if err := s.PrometheusSend(content, len); err != nil {
			klog.Errorf("RobotGroupMsg-PrometheusSend is error", err.Error())
			return err
		}
	}

	return nil
}
//啦啦啦

func (s *Service) PrometheusSend(msgContent string, len int) error {
	msg := http2.SendMsg{
		Msgtype: "markdown",
	}
	key := s.key
	msgContent += "\n" + "还有" + strconv.Itoa(len-1) + "条数没查看"
	msg.Markdown = &http2.MsgContent{Content: msgContent}
	err := http2.DoPost(WechatUrl+key, msg)
	if err != nil {
		klog.Errorf("Webhook:send url false: %v", err.Error())
		return err
	}
	return nil
}
