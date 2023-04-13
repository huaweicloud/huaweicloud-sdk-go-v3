package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 过滤条件
type ReportFilter struct {

	// 创建人
	CreatorIds *string `json:"creatorIds,omitempty"`

	// 所属人
	OwnerIds *string `json:"ownerIds,omitempty"`

	// 级别
	Ranks *string `json:"ranks,omitempty"`

	// releaseId
	ReleaseIds *string `json:"releaseIds,omitempty"`

	// 状态
	Status *string `json:"status,omitempty"`

	// 级别
	ModuleIds *string `json:"moduleIds,omitempty"`

	// 结果
	Results *string `json:"results,omitempty"`

	// 标签
	LabelIds *string `json:"labelIds,omitempty"`

	// 开始时间
	StartTime *string `json:"startTime,omitempty"`

	// 结束时间
	EndTime *string `json:"endTime,omitempty"`

	// 是否关联需求
	IsAssociateIssue *string `json:"isAssociateIssue,omitempty"`
}

func (o ReportFilter) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ReportFilter struct{}"
	}

	return strings.Join([]string{"ReportFilter", string(data)}, " ")
}
