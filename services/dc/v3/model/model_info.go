package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Info struct {

	// 配额类型
	Type *string `json:"type,omitempty"`

	// 可用的配额数，-1 代表不受限制
	Quota *int64 `json:"quota,omitempty"`

	// 已使用的配额数量
	Used *int64 `json:"used,omitempty"`

	// 用量单位
	Unit *string `json:"unit,omitempty"`
}

func (o Info) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Info struct{}"
	}

	return strings.Join([]string{"Info", string(data)}, " ")
}
