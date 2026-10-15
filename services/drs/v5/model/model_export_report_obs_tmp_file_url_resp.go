package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExportReportObsTmpFileUrlResp 导出文件OBS临时URI对象
type ExportReportObsTmpFileUrlResp struct {

	// 文件名
	FileName *string `json:"file_name,omitempty"`

	// OBS文件临时URI
	ObsTempUri *string `json:"obs_temp_uri,omitempty"`
}

func (o ExportReportObsTmpFileUrlResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportReportObsTmpFileUrlResp struct{}"
	}

	return strings.Join([]string{"ExportReportObsTmpFileUrlResp", string(data)}, " ")
}
