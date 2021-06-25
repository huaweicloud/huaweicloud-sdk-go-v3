package model

import (
	"encoding/json"

	"strings"
)

// 审核结果
type ReviewDetail struct {
	// 置信度，取值介于0与1之间。

	Confidence string `json:"confidence"`
	// 每个检测结果的标签化说明，在politics场景中 label为对应的政治人物信息，在terrorism场景中 label为对应的暴恐元素（枪支、刀具、火灾等） 信息，在porn场景中label为对应的涉黄元素（涉 黄、性感等）信息。

	Label *string `json:"label,omitempty"`
}

func (o ReviewDetail) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ReviewDetail struct{}"
	}

	return strings.Join([]string{"ReviewDetail", string(data)}, " ")
}
