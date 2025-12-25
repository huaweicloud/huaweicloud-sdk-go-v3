package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UploadAttachmentRequest Request Object
type UploadAttachmentRequest struct {

	// **参数解释**: 工作空间ID **取值范围**: 不涉及
	WorkspaceId string `json:"workspace_id"`

	Body *UploadAttachmentRequestBody `json:"body,omitempty" type:"multipart"`
}

func (o UploadAttachmentRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UploadAttachmentRequest struct{}"
	}

	return strings.Join([]string{"UploadAttachmentRequest", string(data)}, " ")
}
