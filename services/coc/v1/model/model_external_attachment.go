package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExternalAttachment 上传成功后的附件信息。
type ExternalAttachment struct {

	// 附件唯一id。
	DocId *string `json:"doc_id,omitempty"`

	// 附件文件名。
	DocName *string `json:"doc_name,omitempty"`
}

func (o ExternalAttachment) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExternalAttachment struct{}"
	}

	return strings.Join([]string{"ExternalAttachment", string(data)}, " ")
}
