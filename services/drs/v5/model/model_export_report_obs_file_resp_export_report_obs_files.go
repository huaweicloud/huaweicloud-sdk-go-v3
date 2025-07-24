package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ExportReportObsFileRespExportReportObsFiles struct {

	// 文件名称。
	FileName *string `json:"file_name,omitempty"`

	// 最后修改时间。
	LastModified *string `json:"last_modified,omitempty"`
}

func (o ExportReportObsFileRespExportReportObsFiles) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportReportObsFileRespExportReportObsFiles struct{}"
	}

	return strings.Join([]string{"ExportReportObsFileRespExportReportObsFiles", string(data)}, " ")
}
