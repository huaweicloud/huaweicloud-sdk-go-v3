package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowCurUserRoleResponse Response Object
type ShowCurUserRoleResponse struct {

	// 用户角色id 成员角色, -1 项目创建者, 3 项目经理, 4 开发人员, 5 测试经理, 6 测试人员, 7 参与者, 8 浏览者, 9 运维经理
	UserRole       *int32 `json:"user_role,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowCurUserRoleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCurUserRoleResponse struct{}"
	}

	return strings.Join([]string{"ShowCurUserRoleResponse", string(data)}, " ")
}
