package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowTopTrendRequest Request Object
type ShowTopTrendRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	// 数据库引擎类型
	EngineType string `json:"engine_type"`

	// 库表对象类型
	ObjectType string `json:"object_type"`

	// 库表对象名称
	ObjectName string `json:"object_name"`

	// 库名
	DatabaseName *string `json:"database_name,omitempty"`

	// 开始时间（Unix timestamp），单位：毫秒
	StartTime int64 `json:"start_time"`

	// 结束时间（Unix timestamp），单位：毫秒
	EndTime int64 `json:"end_time"`

	// 节点ID
	NodeId *string `json:"node_id,omitempty"`
}

func (o ShowTopTrendRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTopTrendRequest struct{}"
	}

	return strings.Join([]string{"ShowTopTrendRequest", string(data)}, " ")
}
