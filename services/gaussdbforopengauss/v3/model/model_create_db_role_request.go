package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateDbRoleRequest Request Object
type CreateDbRoleRequest struct {

	// 语言
	XLanguage *string `json:"X-Language,omitempty"`

	// 实例ID。
	InstanceId string `json:"instance_id"`

	Body *CreateDbRoleRequestBody `json:"body,omitempty"`
}

func (o CreateDbRoleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDbRoleRequest struct{}"
	}

	return strings.Join([]string{"CreateDbRoleRequest", string(data)}, " ")
}
