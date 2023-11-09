package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Quota struct {

	// 资源类型
	Type *string `json:"type,omitempty"`

	// 配额上限
	Quota *int32 `json:"quota,omitempty"`

	// 已使用数
	Used *int32 `json:"used,omitempty"`
}

func (o Quota) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Quota struct{}"
	}

	return strings.Join([]string{"Quota", string(data)}, " ")
}
