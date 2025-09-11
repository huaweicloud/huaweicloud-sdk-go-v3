package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateIpdProjectIssueAttachmentResponse Response Object
type CreateIpdProjectIssueAttachmentResponse struct {

	// 返回信息
	Message *string `json:"message,omitempty"`

	// 返回附件对象
	Result *[]AttachmentVo `json:"result,omitempty"`

	// 返回状态
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateIpdProjectIssueAttachmentResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateIpdProjectIssueAttachmentResponse struct{}"
	}

	return strings.Join([]string{"CreateIpdProjectIssueAttachmentResponse", string(data)}, " ")
}
