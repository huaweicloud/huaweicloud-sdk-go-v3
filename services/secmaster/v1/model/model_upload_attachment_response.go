package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UploadAttachmentResponse Response Object
type UploadAttachmentResponse struct {

	// 错误码
	Code *string `json:"code,omitempty"`

	// 错误信息
	Message *string `json:"message,omitempty"`

	// 请求状态
	Success *bool `json:"success,omitempty"`

	// 请求id
	RequestId *string `json:"request_id,omitempty"`

	Data           *UploadAttachmentInfo `json:"data,omitempty"`
	HttpStatusCode int                   `json:"-"`
}

func (o UploadAttachmentResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UploadAttachmentResponse struct{}"
	}

	return strings.Join([]string{"UploadAttachmentResponse", string(data)}, " ")
}
