package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// VersionSetProperty 开源依赖规则详情
type VersionSetProperty struct {

	// 是否启用
	Enable *bool `json:"enable,omitempty"`

	// 规则列表
	Rules *[]VersionSetRule `json:"rules,omitempty"`
}

func (o VersionSetProperty) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VersionSetProperty struct{}"
	}

	return strings.Join([]string{"VersionSetProperty", string(data)}, " ")
}
