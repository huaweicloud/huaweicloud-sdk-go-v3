package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GetDocumentRequest Request Object
type GetDocumentRequest struct {

	// 作业uuid
	DocumentId string `json:"document_id"`

	// 作业版本号，示例值v1、v2，不传默认查询最新版本
	Version *string `json:"version,omitempty"`

	// 执行的作业类型 枚举：PUBLIC/PRIVATE 默认PRIVATE
	DocumentType *string `json:"document_type,omitempty"`
}

func (o GetDocumentRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetDocumentRequest struct{}"
	}

	return strings.Join([]string{"GetDocumentRequest", string(data)}, " ")
}
