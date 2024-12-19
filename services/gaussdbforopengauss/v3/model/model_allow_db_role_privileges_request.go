package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AllowDbRolePrivilegesRequest Request Object
type AllowDbRolePrivilegesRequest struct {

	// 语言
	XLanguage *string `json:"X-Language,omitempty"`

	// 实例ID。
	InstanceId string `json:"instance_id"`

	Body *AllowDbRolePrivilegesRequestBody `json:"body,omitempty"`
}

func (o AllowDbRolePrivilegesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AllowDbRolePrivilegesRequest struct{}"
	}

	return strings.Join([]string{"AllowDbRolePrivilegesRequest", string(data)}, " ")
}
