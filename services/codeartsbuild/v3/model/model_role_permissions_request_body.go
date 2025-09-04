package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RolePermissionsRequestBody 创建更新分组请求体
type RolePermissionsRequestBody struct {

	// CodeArts项目ID。获取方式请参考[获取CodeArts项目ID](https://support.huaweicloud.com/api-codeci/cloudbuild_03_0022.html)。
	ProjectId *string `json:"project_id,omitempty"`

	// 任务id集合
	JobIds *[]string `json:"job_ids,omitempty"`

	// 是否同步最新项目权限
	ProjectSwitch *bool `json:"project_switch,omitempty"`

	// 角色权限信息
	Permissions *[]JobRolePermission `json:"permissions,omitempty"`
}

func (o RolePermissionsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RolePermissionsRequestBody struct{}"
	}

	return strings.Join([]string{"RolePermissionsRequestBody", string(data)}, " ")
}
