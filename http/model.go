package http2

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
>开始于:<font color=\"comment\">{{.StartsAt}}</font>
>Labels:
{{ range $key, $value := .Labels }}
	{{ $key }}:{{ $value }}
{{end}}
>Annotations:
{{ range $key, $value := .Annotations }}
	{{ $key }}:{{ $value }}
{{end}}
>详情:[点击查看]({{.GeneratorURL}})`
