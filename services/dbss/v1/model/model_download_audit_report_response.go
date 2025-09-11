package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DownloadAuditReportResponse Response Object
type DownloadAuditReportResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DownloadAuditReportResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DownloadAuditReportResponse struct{}"
	}

	return strings.Join([]string{"DownloadAuditReportResponse", string(data)}, " ")
}
