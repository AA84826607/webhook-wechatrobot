package http2

import "github.com/prometheus/alertmanager/template"

// wechat
type SendMsg struct {
	Msgtype  string      `json:"msgtype"`
	Markdown interface{} `json:"markdown"`
}

type MsgContent struct {
	Content string `json:"content"`
}

const Templ = `Promethues Alert:
>状态:<font color=\"comment\">{{.Status}}</font>
{{ range $key, $value := .Labels }}
	{{ $value }}-
{{end}}`

func RobotMsgModel(proSend template.Alert)string{
	var data string
	for key,value := range proSend.Labels {
		data+=key+":"+value+"\n"
	}
	return data
}