package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAttachmentRequest Request Object
type ShowAttachmentRequest struct {

	// **参数解释**: 工作空间ID **取值范围**: 不涉及
	WorkspaceId string `json:"workspace_id"`

	// 附件ID
	AttachId string `json:"attach_id"`
}

func (o ShowAttachmentRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAttachmentRequest struct{}"
	}

	return strings.Join([]string{"ShowAttachmentRequest", string(data)}, " ")
}
