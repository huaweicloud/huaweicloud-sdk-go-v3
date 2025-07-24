package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTemplatesResponse Response Object
type ListTemplatesResponse struct {

	// 总数。
	Count *int32 `json:"count,omitempty"`

	// 文件列表。
	ExportReportObsFiles *[]ExportReportObsFileRespExportReportObsFiles `json:"export_report_obs_files,omitempty"`
	HttpStatusCode       int                                            `json:"-"`
}

func (o ListTemplatesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTemplatesResponse struct{}"
	}

	return strings.Join([]string{"ListTemplatesResponse", string(data)}, " ")
}
