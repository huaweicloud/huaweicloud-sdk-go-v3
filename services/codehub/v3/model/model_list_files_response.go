package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListFilesResponse struct {
	Error *Error `json:"error,omitempty" xml:"error"`

	Result *FilesResponseInfo `json:"result,omitempty" xml:"result"`

	// 响应状态
	Status         *string `json:"status,omitempty" xml:"status"`
	HttpStatusCode int     `json:"-"`
}

func (o ListFilesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFilesResponse struct{}"
	}

	return strings.Join([]string{"ListFilesResponse", string(data)}, " ")
}
