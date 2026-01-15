package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAuditTrendHistoryResponse Response Object
type ListAuditTrendHistoryResponse struct {

	// 趋势统计数据列表
	Body           *[]TrendStatusResponse `json:"body,omitempty"`
	HttpStatusCode int                    `json:"-"`
}

func (o ListAuditTrendHistoryResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAuditTrendHistoryResponse struct{}"
	}

	return strings.Join([]string{"ListAuditTrendHistoryResponse", string(data)}, " ")
}
