package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UploadAttachmentInfo 上传时返回的附件内容
type UploadAttachmentInfo struct {

	// 附件id
	Id *string `json:"id,omitempty"`

	// 附件名称
	FileName *string `json:"file_name,omitempty"`

	// 当前的工作空间id
	WorkspaceId *string `json:"workspace_id,omitempty"`
}

func (o UploadAttachmentInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UploadAttachmentInfo struct{}"
	}

	return strings.Join([]string{"UploadAttachmentInfo", string(data)}, " ")
}
