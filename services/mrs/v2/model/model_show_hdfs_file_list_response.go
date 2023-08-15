package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowHdfsFileListResponse Response Object
type ShowHdfsFileListResponse struct {

	// 文件总数，与分页无关。
	TotalCount *int64 `json:"total_count,omitempty"`

	// 文件列表。
	Files          *[]FileStatusV2 `json:"files,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o ShowHdfsFileListResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowHdfsFileListResponse struct{}"
	}

	return strings.Join([]string{"ShowHdfsFileListResponse", string(data)}, " ")
}
