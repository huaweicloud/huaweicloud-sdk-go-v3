package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ResourcesInfo struct {

	// 配额类型
	Type *string `json:"type,omitempty"`

	// 已使用配额
	Used *int32 `json:"used,omitempty"`

	// 总配额数
	Quota *int32 `json:"quota,omitempty"`

	// 最小配额值
	Min *int32 `json:"min,omitempty"`
}

func (o ResourcesInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourcesInfo struct{}"
	}

	return strings.Join([]string{"ResourcesInfo", string(data)}, " ")
}
