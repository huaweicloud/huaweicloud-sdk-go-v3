package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AddMemberRequestV4 struct {

	// 租户名称（跨租户添加用户时，填写正确的租户名称，可将未授权的租户主动授权，将用户添加为项目成员）
	DomainName *string `json:"domain_name,omitempty"`

	// 租户id
	DomainId string `json:"domain_id"`

	// '用户在项目中的角色ID' 成员角色, -1 项目创建者, 3 项目经理, 4 开发人员, 5 测试经理, 6 测试人员, 7 参与者, 8 浏览者, 9 运维经理
	RoleId *int32 `json:"role_id,omitempty"`

	// 用户32位uuid
	UserId string `json:"user_id"`
}

func (o AddMemberRequestV4) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddMemberRequestV4 struct{}"
	}

	return strings.Join([]string{"AddMemberRequestV4", string(data)}, " ")
}
