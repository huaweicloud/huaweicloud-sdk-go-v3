package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowInstancePatchItemsRequest Request Object
type ShowInstancePatchItemsRequest struct {

	// 合规性报告id
	InstanceCompliantId string `json:"instance_compliant_id"`

	// 偏移量
	Offset *int32 `json:"offset,omitempty"`

	// 每页数量
	Limit *int32 `json:"limit,omitempty"`

	// 补丁名称
	Title *string `json:"title,omitempty"`

	// 排序 - asc：升序 - desc：降序
	SortDir *string `json:"sort_dir,omitempty"`

	// 排序字段 -installed_time：补丁安装时间
	SortKey *string `json:"sort_key,omitempty"`

	// 补丁状态 INSTALLED：已安装 INSTALLED_OTHER：已安装其他 MISSING：缺失 REJECT：拒绝 FAILED：失败 PENDING_REBOOT：已安装待重启
	PatchStatus *string `json:"patch_status,omitempty"`

	// 分类
	Classification *string `json:"classification,omitempty"`

	// 严重性级别
	SeverityLevel *string `json:"severity_level,omitempty"`

	// 合规性级别
	ComplianceLevel *string `json:"compliance_level,omitempty"`
}

func (o ShowInstancePatchItemsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowInstancePatchItemsRequest struct{}"
	}

	return strings.Join([]string{"ShowInstancePatchItemsRequest", string(data)}, " ")
}
