package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CheckClickHouseTableConfigResponse Response Object
type CheckClickHouseTableConfigResponse struct {

	// 源数据库名。
	SourceDatabaseName *string `json:"source_database_name,omitempty"`

	// 表配置检查结果。
	TableConfigCheckResults *[]ChDatabaseTableConfigCheckResult `json:"table_config_check_results,omitempty"`
	HttpStatusCode          int                                 `json:"-"`
}

func (o CheckClickHouseTableConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckClickHouseTableConfigResponse struct{}"
	}

	return strings.Join([]string{"CheckClickHouseTableConfigResponse", string(data)}, " ")
}
