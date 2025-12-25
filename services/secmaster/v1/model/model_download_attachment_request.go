package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DownloadAttachmentRequest Request Object
type DownloadAttachmentRequest struct {

	// **参数解释**: 工作空间ID **取值范围**: 不涉及
	WorkspaceId string `json:"workspace_id"`

	// 附件ID
	AttachId string `json:"attach_id"`
}

func (o DownloadAttachmentRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DownloadAttachmentRequest struct{}"
	}

	return strings.Join([]string{"DownloadAttachmentRequest", string(data)}, " ")
}
