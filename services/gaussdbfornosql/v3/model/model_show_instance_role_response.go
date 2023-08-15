package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowInstanceRoleResponse Response Object
type ShowInstanceRoleResponse struct {

	// 枚举类型(master、slave)代表实例是主实例还是备实例。
	Role           *string `json:"role,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowInstanceRoleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowInstanceRoleResponse struct{}"
	}

	return strings.Join([]string{"ShowInstanceRoleResponse", string(data)}, " ")
}
