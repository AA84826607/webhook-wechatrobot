package http2

import (
	"bytes"
	"github.com/prometheus/alertmanager/template"
	"k8s.io/klog"
	gotemplate "text/template"
)

// labelsKey
const (
	LabelsKeySeverity="severity"
	LabelsKeyJob="job"
	LabelsKeyCephId="cephid"
	LabelsKeyPoolId="pool_id"
)

// severityValue
const (
	Warnings="warning"
	Errors="error"
)

var DelKeys = []string{LabelsKeyJob,LabelsKeyPoolId,LabelsKeyCephId}

type SendMsg struct {
	Msgtype  string      `json:"msgtype"`
	Markdown interface{} `json:"markdown"`
}

type MsgContent struct {
	Content string `json:"content"`
}

// messsage models
const Templ = `状态:<font color=\"comment\">{{.Status}}</font>
{{ range $key, $value := .Labels }}
	{{ $key }}:{{ $value }}
{{end}}
>详情:[点击查看]( http://8.129.31.137/alertmanager )`


// one
func RobotMsgModelOne(proSend template.Alert,addDes string)(string,error){
	var doc bytes.Buffer
	t, err := gotemplate.New("alert").Parse(Templ)
	if err != nil {
		klog.Errorf("Webhook: initial go template error: %v", err.Error())
		return "",err
	}
	if err := t.Execute(&doc, proSend); err != nil {
		klog.Errorf("Webhook: go template execute error: %v", err.Error())
		return "",err
	}
	return addDes+"!!!\n"+doc.String(),nil
}

