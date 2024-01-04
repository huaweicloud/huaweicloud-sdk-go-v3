package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TenantQuotaUsed 引擎配额详情
type TenantQuotaUsed struct {

	// 配额类型
	Type *string `json:"type,omitempty"`

	// 已使用配额
	Used *int32 `json:"used,omitempty"`

	// 总配额
	Quota *int32 `json:"quota,omitempty"`
}

func (o TenantQuotaUsed) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TenantQuotaUsed struct{}"
	}

	return strings.Join([]string{"TenantQuotaUsed", string(data)}, " ")
}
