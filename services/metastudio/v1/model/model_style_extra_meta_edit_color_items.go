package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type StyleExtraMetaEditColorItems struct {

	// 最小值
	MinValue *[]float32 `json:"min_value,omitempty"`

	// 最大值
	MaxValue *[]float32 `json:"max_value,omitempty"`

	// 展示标签 {\"cn\": \"xxxxx\",\"en\":\"xxxxx\"}
	Label map[string]string `json:"label,omitempty"`
}

func (o StyleExtraMetaEditColorItems) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StyleExtraMetaEditColorItems struct{}"
	}

	return strings.Join([]string{"StyleExtraMetaEditColorItems", string(data)}, " ")
}
