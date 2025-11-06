package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListFilesResponse Response Object
type ListFilesResponse struct {

	// 文件列表
	Body *[]string `json:"body,omitempty"`

	XTotal         *string `json:"X-Total,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListFilesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFilesResponse struct{}"
	}

	return strings.Join([]string{"ListFilesResponse", string(data)}, " ")
}
