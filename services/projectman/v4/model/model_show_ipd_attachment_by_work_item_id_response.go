package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowIpdAttachmentByWorkItemIdResponse Response Object
type ShowIpdAttachmentByWorkItemIdResponse struct {

	// 请求状态
	Status *string `json:"status,omitempty"`

	// 信息
	Message *string `json:"message,omitempty"`

	// 附件列表
	Result         *[]AttachmentEntity `json:"result,omitempty"`
	HttpStatusCode int                 `json:"-"`
}

func (o ShowIpdAttachmentByWorkItemIdResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowIpdAttachmentByWorkItemIdResponse struct{}"
	}

	return strings.Join([]string{"ShowIpdAttachmentByWorkItemIdResponse", string(data)}, " ")
}
