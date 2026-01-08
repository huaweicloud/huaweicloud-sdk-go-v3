package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExportUsersTemplateRequest Request Object
type ExportUsersTemplateRequest struct {

	// 语言 * zh_CN：中文。 * en_US：英文。
	Language *string `json:"language,omitempty"`

	// 激活类型，默认为用户激活。 * USER_ACTIVATE： 用户激活 * ADMIN_ACTIVATE： 管理员激活
	ActiveType *string `json:"active_type,omitempty"`
}

func (o ExportUsersTemplateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportUsersTemplateRequest struct{}"
	}

	return strings.Join([]string{"ExportUsersTemplateRequest", string(data)}, " ")
}
