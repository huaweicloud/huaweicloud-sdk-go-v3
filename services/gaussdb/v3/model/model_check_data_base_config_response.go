package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CheckDataBaseConfigResponse Response Object
type CheckDataBaseConfigResponse struct {

	// GaussDB(for MySQL)数据库名称。
	SourceDatabaseName *string `json:"source_database_name,omitempty"`

	// GaussDB(for MySQL)数据库配置检查结果。
	SourceDbConfigCheckResults *[]DbConfigCheckResult `json:"source_db_config_check_results,omitempty"`

	// 目标数据库名称。
	TargetDatabaseName *string `json:"target_database_name,omitempty"`

	// 目标数据库配置检查结果。
	TargetDbConfigCheckResults *[]DbConfigCheckResult `json:"target_db_config_check_results,omitempty"`

	// 同步任务名称。
	TaskName       *string `json:"task_name,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CheckDataBaseConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckDataBaseConfigResponse struct{}"
	}

	return strings.Join([]string{"CheckDataBaseConfigResponse", string(data)}, " ")
}
