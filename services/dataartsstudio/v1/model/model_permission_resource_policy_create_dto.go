package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PermissionResourcePolicyCreateDto struct {

	// 策略名称：英文和汉字开头, 支持英文、汉字、数字、下划线, 2-64字符
	PolicyName string `json:"policy_name"`

	// 资源对象列表。资源对象包含 数据连接, 连接获取方法详见[查询数据连接列表](ListDataconnections.xml)
	Resources []ResourcePolicyItem `json:"resources"`

	// 成员列表。 成员包含空间用户、空间用户组、空间用户角色。空间用户、用户组获取方法请参见[获取工作空间用户信息](ListWorkspaceusers.xml),空间角色获取方法参见[获取工作空间用户角色](ListWorkspaceRoles.xml)
	Members []MemberPolicyItem `json:"members"`
}

func (o PermissionResourcePolicyCreateDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PermissionResourcePolicyCreateDto struct{}"
	}

	return strings.Join([]string{"PermissionResourcePolicyCreateDto", string(data)}, " ")
}
