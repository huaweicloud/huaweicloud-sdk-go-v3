package model

import (
	"encoding/json"

	"strings"
)

// 调用成功时表示调用结果。  调用失败时无此字段。
type TextDetectionResponseResult struct {
	// 检测结果是否通过。  block：包含敏感信息，不通过。  pass：不包含敏感信息，通过。  review：需要人工复查。

	Suggestion *string `json:"suggestion,omitempty"`
	// 返回的相关检测结果详细信息：  - politics：涉政敏感词列表。  - porn：涉黄敏感词列表。  - ad：广告敏感词列表。  - abuse：辱骂敏感词列表。  - contraband：违禁品敏感词列表。  - flood：灌水文本。  > - 灌水文本最多显示200个字符。 > - 每个场景的返回结果的详细信息是指命中词，不是返回全部文本。

	Detail *interface{} `json:"detail,omitempty"`
}

func (o TextDetectionResponseResult) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "TextDetectionResponseResult struct{}"
	}

	return strings.Join([]string{"TextDetectionResponseResult", string(data)}, " ")
}
