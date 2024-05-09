package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CheckTableConfigResponse Response Object
type CheckTableConfigResponse struct {

	// GaussDB(for MySQL)数据库名。
	SourceDatabaseName *string `json:"source_database_name,omitempty"`

	// 目标数据库名。
	TargetDatabaseName *string `json:"target_database_name,omitempty"`

	// 表配置检查结果。
	TblConfigCheckResults *[]TableConfigCheckResult `json:"tbl_config_check_results,omitempty"`

	// 同步任务名称。
	TaskName       *string `json:"task_name,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CheckTableConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckTableConfigResponse struct{}"
	}

	return strings.Join([]string{"CheckTableConfigResponse", string(data)}, " ")
}
