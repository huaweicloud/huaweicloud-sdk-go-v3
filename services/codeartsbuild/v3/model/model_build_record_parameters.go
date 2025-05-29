package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BuildRecordParameters struct {

	// 参数名
	Name *string `json:"name,omitempty"`

	// 是否为私密参数
	Secret *bool `json:"secret,omitempty"`

	// 参数值
	Value *string `json:"value,omitempty"`

	// 类型
	Type *string `json:"type,omitempty"`
}

func (o BuildRecordParameters) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BuildRecordParameters struct{}"
	}

	return strings.Join([]string{"BuildRecordParameters", string(data)}, " ")
}
