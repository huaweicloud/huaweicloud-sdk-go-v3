package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TagMatches struct {

	// 为要匹配的字段。 约束：值只能为resource_name。
	Key *string `json:"key,omitempty"`

	// 模糊匹配的值。  约束：最大长度为255个字符，为空返回空值。
	Value *string `json:"value,omitempty"`
}

func (o TagMatches) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TagMatches struct{}"
	}

	return strings.Join([]string{"TagMatches", string(data)}, " ")
}
