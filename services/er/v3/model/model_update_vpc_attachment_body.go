package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 更新vpc基本信息请求体
type UpdateVpcAttachmentBody struct {

	// VPC连接描述信息，取值范围：最大长度36字节，带“-”连字符的UUID格式
	Description *string `json:"description,omitempty"`

	// VPC连接名称，取值范围：1-64个字符，支持数字、字母、中文、_(下划线)、-（中划线）、.（点）
	Name *string `json:"name,omitempty"`
}

func (o UpdateVpcAttachmentBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateVpcAttachmentBody struct{}"
	}

	return strings.Join([]string{"UpdateVpcAttachmentBody", string(data)}, " ")
}
