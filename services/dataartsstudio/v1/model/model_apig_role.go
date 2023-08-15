package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ApigRole IAM用户组信息
type ApigRole struct {

	// 角色id，r00001：管理员；r00002：开发者；r00003：运维者；r00004：访客
	RoleId *string `json:"role_id,omitempty"`
}

func (o ApigRole) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApigRole struct{}"
	}

	return strings.Join([]string{"ApigRole", string(data)}, " ")
}
