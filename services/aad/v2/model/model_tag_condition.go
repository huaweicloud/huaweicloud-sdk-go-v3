package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TagCondition struct {

	// 防护动作
	Category *string `json:"category,omitempty"`

	// 字段内容
	Contents *[]string `json:"contents,omitempty"`
}

func (o TagCondition) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TagCondition struct{}"
	}

	return strings.Join([]string{"TagCondition", string(data)}, " ")
}
