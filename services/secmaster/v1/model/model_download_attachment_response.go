package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DownloadAttachmentResponse Response Object
type DownloadAttachmentResponse struct {

	// 错误码
	Code *string `json:"code,omitempty"`

	// 错误信息
	Message *string `json:"message,omitempty"`

	// 请求状态
	Success *bool `json:"success,omitempty"`

	// 请求id
	RequestId *string `json:"request_id,omitempty"`

	// 附件的obs下载地址
	Data           *string `json:"data,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DownloadAttachmentResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DownloadAttachmentResponse struct{}"
	}

	return strings.Join([]string{"DownloadAttachmentResponse", string(data)}, " ")
}
