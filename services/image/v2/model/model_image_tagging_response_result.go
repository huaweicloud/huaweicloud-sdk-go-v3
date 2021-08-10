package model

import (
	"encoding/json"

	"strings"
)

// 调用成功时为图片标签内容。  调用失败时无此字段。
type ImageTaggingResponseResult struct {
	// 标签列表集合。

	Tags *[]ImageTaggingItemBody `json:"tags,omitempty"`
}

func (o ImageTaggingResponseResult) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ImageTaggingResponseResult struct{}"
	}

	return strings.Join([]string{"ImageTaggingResponseResult", string(data)}, " ")
}
