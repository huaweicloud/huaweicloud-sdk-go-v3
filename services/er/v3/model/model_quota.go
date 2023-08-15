package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Quota struct {

	// 配额类型
	QuotaKey *string `json:"quota_key,omitempty"`

	// 可用的配额数，-1 代表不受限制
	QuotaLimit *int64 `json:"quota_limit,omitempty"`

	// 已使用的配额数量
	Used *int64 `json:"used,omitempty"`

	// 用量单位
	Unit *string `json:"unit,omitempty"`
}

func (o Quota) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Quota struct{}"
	}

	return strings.Join([]string{"Quota", string(data)}, " ")
}
