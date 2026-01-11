# AI Product 

- 支持 Gemini + OpenAI 多模型切换
- Prompt 可资产化，可版本管理
- HTTP /generate 接口可直接对接前端

## 运行

```bash
export OPENAI_API_KEY=你的OpenAIKey
export GEMINI_API_KEY=你的GeminiKey
go run ./cmd/server
