package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AllocateSpResourceSummaryInfo 分配给租户的资源概览。
type AllocateSpResourceSummaryInfo struct {

	// 资源规格编码
	ResourceSpecCode *string `json:"resource_spec_code,omitempty"`

	// 资源数量。
	ResourceNum *float64 `json:"resource_num,omitempty"`

	// 资源已使用数量。
	ResourceUsedNum *float64 `json:"resource_used_num,omitempty"`

	// 资源在TMS上的剩余量
	ResourceRemindQuota *float64 `json:"resource_remind_quota,omitempty"`
}

func (o AllocateSpResourceSummaryInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AllocateSpResourceSummaryInfo struct{}"
	}

	return strings.Join([]string{"AllocateSpResourceSummaryInfo", string(data)}, " ")
}
