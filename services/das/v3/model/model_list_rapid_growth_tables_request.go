package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListRapidGrowthTablesRequest Request Object
type ListRapidGrowthTablesRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	// 数据库引擎类型
	EngineType string `json:"engine_type"`

	// 数据库名称
	DatabaseName *string `json:"database_name,omitempty"`

	// 关键字
	Keyword *string `json:"keyword,omitempty"`
}

func (o ListRapidGrowthTablesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRapidGrowthTablesRequest struct{}"
	}

	return strings.Join([]string{"ListRapidGrowthTablesRequest", string(data)}, " ")
}
