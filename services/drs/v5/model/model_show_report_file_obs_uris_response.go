package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowReportFileObsUrisResponse Response Object
type ShowReportFileObsUrisResponse struct {

	// 可下载文件总数
	Count *int32 `json:"count,omitempty"`

	// 可下载的文件列表
	ExportReportObsFiles *[]ExportReportObsTmpFileUrlResp `json:"export_report_obs_files,omitempty"`
	HttpStatusCode       int                              `json:"-"`
}

func (o ShowReportFileObsUrisResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowReportFileObsUrisResponse struct{}"
	}

	return strings.Join([]string{"ShowReportFileObsUrisResponse", string(data)}, " ")
}
