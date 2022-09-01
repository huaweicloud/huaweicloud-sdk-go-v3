package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type UploadExtensionFileResponse struct {
	Error *Error `json:"error,omitempty" xml:"error"`

	// 结果
	Result *interface{} `json:"result,omitempty" xml:"result"`

	// 状态
	Status         *string `json:"status,omitempty" xml:"status"`
	HttpStatusCode int     `json:"-"`
}

func (o UploadExtensionFileResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UploadExtensionFileResponse struct{}"
	}

	return strings.Join([]string{"UploadExtensionFileResponse", string(data)}, " ")
}
