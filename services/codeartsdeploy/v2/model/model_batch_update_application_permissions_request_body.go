package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchUpdateApplicationPermissionsRequestBody struct {

	// 项目id
	ProjectId string `json:"project_id"`

	// 应用列表
	ApplicationIds []string `json:"application_ids"`

	// 角色权限
	Roles []AppPermission `json:"roles"`
}

func (o BatchUpdateApplicationPermissionsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUpdateApplicationPermissionsRequestBody struct{}"
	}

	return strings.Join([]string{"BatchUpdateApplicationPermissionsRequestBody", string(data)}, " ")
}
