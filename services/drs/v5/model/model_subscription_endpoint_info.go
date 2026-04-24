package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SubscriptionEndpointInfo struct {

	// 数据库实例id
	DbInstanceId string `json:"db_instance_id"`

	// 数据库名称
	Name string `json:"name"`

	// 数据库内网ip
	Ip string `json:"ip"`

	// 数据库端口
	Port int32 `json:"port"`

	// 数据库类型
	Type string `json:"type"`

	// 数据库用户名
	UserName string `json:"user_name"`
}

func (o SubscriptionEndpointInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SubscriptionEndpointInfo struct{}"
	}

	return strings.Join([]string{"SubscriptionEndpointInfo", string(data)}, " ")
}
