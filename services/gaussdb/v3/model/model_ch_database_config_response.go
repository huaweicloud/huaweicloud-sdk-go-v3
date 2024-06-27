package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChDatabaseConfigResponse 库配置校验成功返回值。
type ChDatabaseConfigResponse struct {

	// 源数据库名称。
	DatabaseName string `json:"database_name"`

	// 源数据库配置检查结果。
	DbConfigCheckResults []ChDatabaseConfigCheckResult `json:"db_config_check_results"`
}

func (o ChDatabaseConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChDatabaseConfigResponse struct{}"
	}

	return strings.Join([]string{"ChDatabaseConfigResponse", string(data)}, " ")
}
