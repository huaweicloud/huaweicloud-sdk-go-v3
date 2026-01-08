package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DownloadStepImageNewRequest Request Object
type DownloadStepImageNewRequest struct {

	// 项目id
	ProjectId string `json:"project_id"`

	// 父级目录名称
	Parent string `json:"parent"`

	// 子级目录名称
	Sub string `json:"sub"`

	// 文件名
	FileName string `json:"file_name"`

	// 文件类型
	FileType string `json:"file_type"`
}

func (o DownloadStepImageNewRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DownloadStepImageNewRequest struct{}"
	}

	return strings.Join([]string{"DownloadStepImageNewRequest", string(data)}, " ")
}
