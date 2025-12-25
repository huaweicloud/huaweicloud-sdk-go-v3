package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UploadAgentFileResponse Response Object
type UploadAgentFileResponse struct {

	// 临时url
	Url *string `json:"url,omitempty"`

	// headers
	Headers map[string]string `json:"headers,omitempty"`

	// 文件名
	FileName       *string `json:"file_name,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UploadAgentFileResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UploadAgentFileResponse struct{}"
	}

	return strings.Join([]string{"UploadAgentFileResponse", string(data)}, " ")
}
