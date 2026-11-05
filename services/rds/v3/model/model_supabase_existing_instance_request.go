package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SupabaseExistingInstanceRequest 存量PG实例请求（manage_type=existed_instance时必填）。
type SupabaseExistingInstanceRequest struct {

	// 数据库实例ID（必须为PostgreSQL实例）。
	InstanceId string `json:"instance_id"`

	// 对接的数据库用户名。
	DatabaseName string `json:"database_name"`

	// 对接的数据库密码。
	DatabasePwd string `json:"database_pwd"`
}

func (o SupabaseExistingInstanceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SupabaseExistingInstanceRequest struct{}"
	}

	return strings.Join([]string{"SupabaseExistingInstanceRequest", string(data)}, " ")
}
