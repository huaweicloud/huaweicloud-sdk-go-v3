package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListBuildInfoRecordBodyResult 结果
type ListBuildInfoRecordBodyResult struct {

	// 健康度
	HealthScore *string `json:"health_score,omitempty"`

	// 分页页数
	PageIndex *string `json:"page_index,omitempty"`

	// 总页数
	TotalPage *string `json:"total_page,omitempty"`

	// 总条数
	TotalRecord *string `json:"total_record,omitempty"`

	// 构建历史详情列表
	JobBuildStates *[]BuildInfoRecord `json:"job_build_states,omitempty"`
}

func (o ListBuildInfoRecordBodyResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBuildInfoRecordBodyResult struct{}"
	}

	return strings.Join([]string{"ListBuildInfoRecordBodyResult", string(data)}, " ")
}
