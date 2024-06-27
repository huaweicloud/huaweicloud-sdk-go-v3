package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CheckClickHouseDataBaseConfigResponse Response Object
type CheckClickHouseDataBaseConfigResponse struct {

	// 源数据库名称。
	DatabaseName *string `json:"database_name,omitempty"`

	// 源数据库配置检查结果。
	DbConfigCheckResults *[]ChDatabaseConfigCheckResult `json:"db_config_check_results,omitempty"`
	HttpStatusCode       int                            `json:"-"`
}

func (o CheckClickHouseDataBaseConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckClickHouseDataBaseConfigResponse struct{}"
	}

	return strings.Join([]string{"CheckClickHouseDataBaseConfigResponse", string(data)}, " ")
}
