package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowFilesResponse Response Object
type ShowFilesResponse struct {

	// 下载文件总数
	TotalCount *int64 `json:"total_count,omitempty"`

	// 下载的文件列表
	Files          *[]FileInfo `json:"files,omitempty"`
	HttpStatusCode int         `json:"-"`
}

func (o ShowFilesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowFilesResponse struct{}"
	}

	return strings.Join([]string{"ShowFilesResponse", string(data)}, " ")
}
