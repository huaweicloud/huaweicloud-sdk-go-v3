package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowClusterRedistributionRequest Request Object
type ShowClusterRedistributionRequest struct {

	// 集群ID
	ClusterId string `json:"cluster_id"`

	// 分页查询，每页大小
	Limit *int32 `json:"limit,omitempty"`

	// 分页查询，偏移
	Offset *int32 `json:"offset,omitempty"`

	// 数据库名称
	DbName *string `json:"db_name,omitempty"`

	// 表名称
	TableName *string `json:"table_name,omitempty"`

	// 类型
	Type *string `json:"type,omitempty"`
}

func (o ShowClusterRedistributionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowClusterRedistributionRequest struct{}"
	}

	return strings.Join([]string{"ShowClusterRedistributionRequest", string(data)}, " ")
}
