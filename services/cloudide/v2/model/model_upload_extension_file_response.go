package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UploadExtensionFileResponse Response Object
type UploadExtensionFileResponse struct {
	Error *Error `json:"error,omitempty"`

	// 结果
	Result *interface{} `json:"result,omitempty"`

	// 状态
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UploadExtensionFileResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UploadExtensionFileResponse struct{}"
	}

	return strings.Join([]string{"UploadExtensionFileResponse", string(data)}, " ")
}
