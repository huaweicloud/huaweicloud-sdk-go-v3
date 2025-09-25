package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowImageFilesStatResponse Response Object
type ShowImageFilesStatResponse struct {

	// 镜像文件总数
	TotalFilesNum *int32 `json:"total_files_num,omitempty"`

	// 镜像文件大小
	TotalFilesSize *int32 `json:"total_files_size,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowImageFilesStatResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowImageFilesStatResponse struct{}"
	}

	return strings.Join([]string{"ShowImageFilesStatResponse", string(data)}, " ")
}
