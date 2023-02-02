package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// update-stack request
type UpdateStackRequestBody struct {

	// 资源栈Id
	StackId *string `json:"stack_id,omitempty"`

	// 资源栈的描述
	Description *string `json:"description,omitempty"`

	// 是否开启删除保护
	EnableDeletionProtection *bool `json:"enable_deletion_protection,omitempty"`

	// 是否开启自动回滚
	EnableAutoRollback *bool `json:"enable_auto_rollback,omitempty"`

	// 委托列表
	Agencies *[]Agency `json:"agencies,omitempty"`
}

func (o UpdateStackRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateStackRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateStackRequestBody", string(data)}, " ")
}
