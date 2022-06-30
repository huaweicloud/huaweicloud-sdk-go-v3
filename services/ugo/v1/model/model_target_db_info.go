package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 目标数据库信息。
type TargetDbInfo struct {

	// 用户名。
	UserName string `json:"user_name"`

	// 用户密码。
	Password string `json:"password"`

	// service名称。
	ServiceName string `json:"service_name"`

	// RDS数据库的实例ID。
	InstanceId string `json:"instance_id"`
}

func (o TargetDbInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TargetDbInfo struct{}"
	}

	return strings.Join([]string{"TargetDbInfo", string(data)}, " ")
}
