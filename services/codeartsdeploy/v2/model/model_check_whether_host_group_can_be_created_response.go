package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CheckWhetherHostGroupCanBeCreatedResponse Response Object
type CheckWhetherHostGroupCanBeCreatedResponse struct {

	// 是否有创建主机集群权限，true 有权限 false 无权限
	CanCreated     *bool `json:"can_created,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o CheckWhetherHostGroupCanBeCreatedResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckWhetherHostGroupCanBeCreatedResponse struct{}"
	}

	return strings.Join([]string{"CheckWhetherHostGroupCanBeCreatedResponse", string(data)}, " ")
}
