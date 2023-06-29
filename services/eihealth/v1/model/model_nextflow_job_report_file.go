package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NextflowJobReportFile 作业report文件
type NextflowJobReportFile struct {

	// 报告文件名
	Name string `json:"name"`

	// 报告文件下载地址
	DownloadUrl string `json:"download_url"`
}

func (o NextflowJobReportFile) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NextflowJobReportFile struct{}"
	}

	return strings.Join([]string{"NextflowJobReportFile", string(data)}, " ")
}
