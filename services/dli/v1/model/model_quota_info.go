package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type QuotaInfo struct {
	Type *string `json:"type,omitempty"`

	// 配额的最小值
	Min *int32 `json:"min,omitempty"`

	// 配额的最大值
	Max *int32 `json:"max,omitempty"`

	// 目前的配额
	Quota *int32 `json:"quota,omitempty"`

	// 已用的配额
	Used *int32 `json:"used,omitempty"`
}

func (o QuotaInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QuotaInfo struct{}"
	}

	return strings.Join([]string{"QuotaInfo", string(data)}, " ")
}
