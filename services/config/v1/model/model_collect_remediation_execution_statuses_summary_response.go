package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CollectRemediationExecutionStatusesSummaryResponse Response Object
type CollectRemediationExecutionStatusesSummaryResponse struct {

	// 合规规则修正执行结果查询列表。
	Value *[]RemediationExecutionStatus `json:"value,omitempty"`

	PageInfo       *PageInfo `json:"page_info,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o CollectRemediationExecutionStatusesSummaryResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CollectRemediationExecutionStatusesSummaryResponse struct{}"
	}

	return strings.Join([]string{"CollectRemediationExecutionStatusesSummaryResponse", string(data)}, " ")
}
