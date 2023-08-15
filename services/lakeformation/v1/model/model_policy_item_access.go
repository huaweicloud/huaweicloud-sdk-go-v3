package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PolicyItemAccess struct {

	// 是否允许
	IsAllowed *bool `json:"is_allowed,omitempty"`

	// 类型
	Type *string `json:"type,omitempty"`
}

func (o PolicyItemAccess) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PolicyItemAccess struct{}"
	}

	return strings.Join([]string{"PolicyItemAccess", string(data)}, " ")
}
