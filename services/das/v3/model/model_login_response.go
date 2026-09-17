package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// LoginResponse Response Object
type LoginResponse struct {

	// 连接ID
	ConnectionId *string `json:"connection_id,omitempty"`

	// 实例名
	InstanceName *string `json:"instance_name,omitempty"`

	// 实例ID
	InstanceId *string `json:"instance_id,omitempty"`

	// 登录名
	LoginUser *string `json:"login_user,omitempty"`

	// 登录的数据库名
	DatabaseName *string `json:"database_name,omitempty"`

	// 引擎类型
	EngineType     *string `json:"engine_type,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o LoginResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LoginResponse struct{}"
	}

	return strings.Join([]string{"LoginResponse", string(data)}, " ")
}
