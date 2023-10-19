package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateRolesRequest Request Object
type UpdateRolesRequest struct {

	// LakeFormation实例ID。创建实例时自动生成。例如：2180518f-42b8-4947-b20b-adfc53981a25。
	InstanceId string `json:"instance_id"`

	// 用户名。只能包含字母、数字、下划线和中划线，且长度为1~256个字符。
	UserName string `json:"user_name"`

	Body *[]RoleInfoInput `json:"body,omitempty"`
}

func (o UpdateRolesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateRolesRequest struct{}"
	}

	return strings.Join([]string{"UpdateRolesRequest", string(data)}, " ")
}
