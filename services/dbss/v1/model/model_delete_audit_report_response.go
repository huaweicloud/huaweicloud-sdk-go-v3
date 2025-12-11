package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteAuditReportResponse Response Object
type DeleteAuditReportResponse struct {

	// 状态  - SUCCESS: 成功  - FAILED: 失败
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteAuditReportResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAuditReportResponse struct{}"
	}

	return strings.Join([]string{"DeleteAuditReportResponse", string(data)}, " ")
}
