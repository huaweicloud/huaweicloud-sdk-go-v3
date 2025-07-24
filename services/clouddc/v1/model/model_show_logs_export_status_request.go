package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowLogsExportStatusRequest Request Object
type ShowLogsExportStatusRequest struct {

	// imetal server id
	Id string `json:"id"`

	// 日志导出id
	ExportId string `json:"export_id"`
}

func (o ShowLogsExportStatusRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowLogsExportStatusRequest struct{}"
	}

	return strings.Join([]string{"ShowLogsExportStatusRequest", string(data)}, " ")
}
