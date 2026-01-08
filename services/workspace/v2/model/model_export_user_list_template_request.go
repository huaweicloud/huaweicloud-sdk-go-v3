package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExportUserListTemplateRequest Request Object
type ExportUserListTemplateRequest struct {

	// 语言 - zh_CN：中文。 - en_US：英文。
	Language *string `json:"language,omitempty"`

	// 系统参数。
	OsType *string `json:"os_type,omitempty"`
}

func (o ExportUserListTemplateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportUserListTemplateRequest struct{}"
	}

	return strings.Join([]string{"ExportUserListTemplateRequest", string(data)}, " ")
}
