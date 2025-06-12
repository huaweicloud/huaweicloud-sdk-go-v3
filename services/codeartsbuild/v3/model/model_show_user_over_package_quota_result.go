package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowUserOverPackageQuotaResult 结果
type ShowUserOverPackageQuotaResult struct {

	// 套餐超期
	IsOverQuota *bool `json:"is_over_quota,omitempty"`

	// 构建套餐
	BuildQuota *string `json:"build_quota,omitempty"`

	// 已使用的套餐用量，套餐为unlimit时无此信息
	UsedQuota *string `json:"used_quota,omitempty"`
}

func (o ShowUserOverPackageQuotaResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowUserOverPackageQuotaResult struct{}"
	}

	return strings.Join([]string{"ShowUserOverPackageQuotaResult", string(data)}, " ")
}
