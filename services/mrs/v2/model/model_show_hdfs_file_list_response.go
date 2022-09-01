package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowHdfsFileListResponse struct {

	// 文件总数，与分页无关。
	TotalCount *int64 `json:"total_count,omitempty" xml:"total_count"`

	// 文件列表。
	Files          *[]FileStatusV2 `json:"files,omitempty" xml:"files"`
	HttpStatusCode int             `json:"-"`
}

func (o ShowHdfsFileListResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowHdfsFileListResponse struct{}"
	}

	return strings.Join([]string{"ShowHdfsFileListResponse", string(data)}, " ")
}
