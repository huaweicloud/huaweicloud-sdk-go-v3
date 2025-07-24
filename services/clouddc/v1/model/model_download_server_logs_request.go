package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DownloadServerLogsRequest Request Object
type DownloadServerLogsRequest struct {

	// imetal server id
	Id string `json:"id"`

	// 日志导出id
	ExportId string `json:"export_id"`
}

func (o DownloadServerLogsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DownloadServerLogsRequest struct{}"
	}

	return strings.Join([]string{"DownloadServerLogsRequest", string(data)}, " ")
}
