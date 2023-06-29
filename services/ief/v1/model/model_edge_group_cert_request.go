package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EdgeGroupCertRequest 边缘节点组证书基本信息
type EdgeGroupCertRequest struct {

	// 证书名称。只允许中文字符、英文字母、数字、下划线、中划线，最大长度64
	Name *string `json:"name,omitempty"`

	// 证书类型，支持填写： - system：创建节点时会默认创建一套系统证书 - application：应用证书 - device：设备证书
	Type *string `json:"type,omitempty"`

	// 证书描述。最大长度为255个字符
	Description *string `json:"description,omitempty"`
}

func (o EdgeGroupCertRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EdgeGroupCertRequest struct{}"
	}

	return strings.Join([]string{"EdgeGroupCertRequest", string(data)}, " ")
}
