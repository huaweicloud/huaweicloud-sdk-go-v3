package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateMembersRoleV4RequestBody struct {

	// 成员角色, -1 项目创建者, 3 项目经理, 4 开发人员, 5 测试经理, 6 测试人员, 7 参与者, 8 浏览者, 9 运维经理
	RoleId int32 `json:"role_id"`

	// 用户id
	UserIds []string `json:"user_ids"`
}

func (o UpdateMembersRoleV4RequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateMembersRoleV4RequestBody struct{}"
	}

	return strings.Join([]string{"UpdateMembersRoleV4RequestBody", string(data)}, " ")
}
