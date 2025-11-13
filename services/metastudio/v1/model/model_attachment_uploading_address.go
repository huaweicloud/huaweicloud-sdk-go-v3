package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AttachmentUploadingAddress 审核附件上传地址
type AttachmentUploadingAddress struct {

	// 序号
	Number *int32 `json:"number,omitempty"`

	// 上传url
	UploadingUrl *string `json:"uploading_url,omitempty"`
}

func (o AttachmentUploadingAddress) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AttachmentUploadingAddress struct{}"
	}

	return strings.Join([]string{"AttachmentUploadingAddress", string(data)}, " ")
}
