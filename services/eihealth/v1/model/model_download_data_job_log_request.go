package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DownloadDataJobLogRequest Request Object
type DownloadDataJobLogRequest struct {

	// 医疗智能体平台项目ID，您可以在EIHealth平台单击所需的项目名称，进入项目设置页面查看。
	EihealthProjectId string `json:"eihealth_project_id"`

	// 数据作业id
	DataJobId string `json:"data_job_id"`
}

func (o DownloadDataJobLogRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DownloadDataJobLogRequest struct{}"
	}

	return strings.Join([]string{"DownloadDataJobLogRequest", string(data)}, " ")
}
