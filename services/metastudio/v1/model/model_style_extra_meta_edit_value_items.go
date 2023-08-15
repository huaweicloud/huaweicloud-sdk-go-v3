package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type StyleExtraMetaEditValueItems struct {

	// 最小值
	MinValue *float32 `json:"min_value,omitempty"`

	// 最大值
	MaxValue *float32 `json:"max_value,omitempty"`

	// 展示标签 {\"cn\": \"xxxxx\",\"en\":\"xxxxx\"}
	Label map[string]string `json:"label,omitempty"`
}

func (o StyleExtraMetaEditValueItems) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StyleExtraMetaEditValueItems struct{}"
	}

	return strings.Join([]string{"StyleExtraMetaEditValueItems", string(data)}, " ")
}
