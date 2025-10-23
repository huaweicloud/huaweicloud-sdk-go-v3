package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowReportResponse Response Object
type ShowReportResponse struct {

	// 报告下载链接
	ReportDownloadUrl *string `json:"report_download_url,omitempty"`

	ReportInfo     *ReportEntity `json:"report_info,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o ShowReportResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowReportResponse struct{}"
	}

	return strings.Join([]string{"ShowReportResponse", string(data)}, " ")
}
