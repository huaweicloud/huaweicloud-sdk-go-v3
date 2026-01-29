package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type InterActiveParams struct {

	// 键
	Key *string `json:"key,omitempty"`

	// 值
	Value *string `json:"value,omitempty"`

	// 查询操作
	Operator *string `json:"operator,omitempty"`
}

func (o InterActiveParams) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InterActiveParams struct{}"
	}

	return strings.Join([]string{"InterActiveParams", string(data)}, " ")
}
