package http2

import (
	"bytes"
	"github.com/prometheus/alertmanager/template"
	"k8s.io/klog"
	gotemplate "text/template"
)

// labelsKey
const (
	LabelsKeySeverity = "severity"
	LabelsKeyJob      = "job"
	LabelsKeyCephId   = "cephid"
	LabelsKeyPoolId   = "pool_id"
)

// severityValue
const (
	Warnings = "warning"
	Errors   = "error"
)

var DelKeys = []string{LabelsKeyJob, LabelsKeyPoolId, LabelsKeyCephId}

type SendMsg struct {
	Msgtype  string      `json:"msgtype"`
	Markdown interface{} `json:"markdown"`
}

type MsgContent struct {
	Content string `json:"content"`
}

////这是zhuyou的代码 ， 有问题就恢复这里
//// messsage models
//const Templ = `状态:<font color=\"comment\">{{.Status}}</font>
//{{ range $key, $value := .Labels }}
//	{{ $key }}:{{ $value }}
//{{end}}
//>详情:[点击查看]( http://8.129.31.137/alertmanager )`
////const Templ =
////
////
////
////`
////{{ range $key, $value := .Labels }}
////{{$key.instance}}:  扇区类型：{{$key.state}},数量：{{$key.value}}
////{{ end }}
////`
//
//// one
//func RobotMsgModelOne(proSend template.Alert, addDes string) (string, error) {
//	var doc bytes.Buffer
//	t, err := gotemplate.New("alert").Parse(Templ)
//	if err != nil {
//		klog.Errorf("Webhook: initial go template error: %v", err.Error())
//		return "", err
//	}
//	if err := t.Execute(&doc, proSend); err != nil {
//		klog.Errorf("Webhook: go template execute error: %v", err.Error())
//		return "", err
//	}
//	return addDes + "!!!\n" + doc.String(), nil
//}
//啦啦啦

//这是测试ceph和lotus分组的代码
const TemplCeph1 = `状态:<font color=\"comment\">{{.Status}}</font>
{{ range $key, $value := .Labels }}
	{{ $key }}:{{ $value }}
{{end}}
>详情:[点击查看]( http://8.129.31.137/alertmanager )`
const TemplGalaxy1 = `
{{.Labels.miner_host}}:{{.Labels.state}} {{.Labels.value}}
{{end}}
>详情:[点击查看]( http://8.129.31.137/alertmanager )`

// one
func RobotMsgModelOne(proSend template.Alert, addDes string) (string, error) {
	var doc bytes.Buffer
	t, err := gotemplate.New("alert").Parse(TemplCeph1)
	if err != nil {
		klog.Errorf("Webhook: initial go template error: %v", err.Error())
		return "", err
	}
	if err := t.Execute(&doc, proSend); err != nil {
		klog.Errorf("Webhook: go template execute error: %v", err.Error())
		return "", err
	}
	return addDes + "!!!\n" + doc.String(), nil
}
func RobotMsgModelTwo(proSend template.Alert, addDes string) (string, error) {
	var doc bytes.Buffer
	t, err := gotemplate.New("alert").Parse(TemplGalaxy1)
	if err != nil {
		klog.Errorf("Webhook: initial go template error: %v", err.Error())
		return "", err
	}
	if err := t.Execute(&doc, proSend); err != nil {
		klog.Errorf("Webhook: go template execute error: %v", err.Error())
		return "", err
	}
	return addDes + "!!!\n" + doc.String(), nil
}