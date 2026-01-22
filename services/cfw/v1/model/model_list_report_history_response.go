package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListReportHistoryResponse Response Object
type ListReportHistoryResponse struct {
	Data           *ListReportHistoryRespData `json:"data,omitempty"`
	HttpStatusCode int                        `json:"-"`
}

func (o ListReportHistoryResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListReportHistoryResponse struct{}"
	}

	return strings.Join([]string{"ListReportHistoryResponse", string(data)}, " ")
}
