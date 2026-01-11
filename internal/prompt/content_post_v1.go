package prompt

var ContentPostV1 = Prompt{
	Code:    "CONTENT_POST",
	Version: "v1",
	System:  "你是一个专业的内容产品生成引擎，输出结构清晰、可直接发布。",
	UserTemplate: `
主题：{{.topic}}
受众：{{.audience}}

要求：
1. 语言自然，不像 AI
2. 有清晰观点
3. 可直接发布
`,
}
