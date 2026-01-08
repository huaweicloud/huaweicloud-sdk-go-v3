package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExportTerminalsBindingDesktopsTemplateRequest Request Object
type ExportTerminalsBindingDesktopsTemplateRequest struct {

	// 语言。
	Language *string `json:"language,omitempty"`
}

func (o ExportTerminalsBindingDesktopsTemplateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportTerminalsBindingDesktopsTemplateRequest struct{}"
	}

	return strings.Join([]string{"ExportTerminalsBindingDesktopsTemplateRequest", string(data)}, " ")
}
