package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateAttachmentBody 更新连接基本信息
type UpdateAttachmentBody struct {

	// 连接描述信息，取值范围：最大长度36字节，带“-”连字符的UUID格式
	Description *string `json:"description,omitempty"`

	// 连接名称，取值范围：1-64个字符，支持数字、字母、中文、_(下划线)、-（中划线）、.（点）
	Name *string `json:"name,omitempty"`
}

func (o UpdateAttachmentBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAttachmentBody struct{}"
	}

	return strings.Join([]string{"UpdateAttachmentBody", string(data)}, " ")
}
