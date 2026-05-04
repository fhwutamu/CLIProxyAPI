package responses

import (
	"testing"

	"github.com/tidwall/gjson"
)

func TestConvertOpenAIResponsesRequestToGemini_UsesUserRoleForFunctionOutput(t *testing.T) {
	input := []byte(`{
		"model": "gemini-3.1-pro-preview",
		"input": [
			{
				"type": "function_call",
				"call_id": "call_123",
				"name": "lookup",
				"arguments": "{\"query\":\"ok\"}"
			},
			{
				"type": "function_call_output",
				"call_id": "call_123",
				"output": "{\"result\":\"done\"}"
			}
		]
	}`)

	out := ConvertOpenAIResponsesRequestToGemini("gemini-3.1-pro-preview", input, false)

	if got := gjson.GetBytes(out, "contents.1.role").String(); got != "user" {
		t.Fatalf("function_call_output role = %q, want user; body=%s", got, string(out))
	}
	if !gjson.GetBytes(out, "contents.1.parts.0.functionResponse").Exists() {
		t.Fatalf("missing functionResponse part; body=%s", string(out))
	}
	if gjson.GetBytes(out, "contents.0.parts.0.functionCall.id").Exists() {
		t.Fatalf("functionCall id must be omitted for Vertex compatibility; body=%s", string(out))
	}
	if gjson.GetBytes(out, "contents.1.parts.0.functionResponse.id").Exists() {
		t.Fatalf("functionResponse id must be omitted for Vertex compatibility; body=%s", string(out))
	}
}
