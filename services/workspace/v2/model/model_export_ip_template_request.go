package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExportIpTemplateRequest Request Object
type ExportIpTemplateRequest struct {

	// 语言 * zh_CN：中文。 * en_US：英文。
	Language *string `json:"language,omitempty"`
}

func (o ExportIpTemplateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportIpTemplateRequest struct{}"
	}

	return strings.Join([]string{"ExportIpTemplateRequest", string(data)}, " ")
}
