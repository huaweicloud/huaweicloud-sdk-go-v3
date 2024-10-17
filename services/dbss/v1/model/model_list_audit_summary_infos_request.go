package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAuditSummaryInfosRequest Request Object
type ListAuditSummaryInfosRequest struct {

	// 偏移量
	Offset *string `json:"offset,omitempty"`

	// 查询记录数
	Limit *string `json:"limit,omitempty"`
}

func (o ListAuditSummaryInfosRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAuditSummaryInfosRequest struct{}"
	}

	return strings.Join([]string{"ListAuditSummaryInfosRequest", string(data)}, " ")
}
