package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowTopDataRequest Request Object
type ShowTopDataRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	// 数据库引擎类型
	EngineType string `json:"engine_type"`

	// 库表对象类型
	ObjectType string `json:"object_type"`

	// 结束时间（Unix timestamp），单位：毫秒
	EndTime int64 `json:"end_time"`

	// 节点ID
	NodeId *string `json:"node_id,omitempty"`

	// 排序字段
	OrderBy *string `json:"order_by,omitempty"`

	// 排序方式
	Order *string `json:"order,omitempty"`

	// 关键字
	Keyword *string `json:"keyword,omitempty"`

	// 页数
	PageNum *int32 `json:"page_num,omitempty"`

	// 页大小
	PageSize *int32 `json:"page_size,omitempty"`
}

func (o ShowTopDataRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTopDataRequest struct{}"
	}

	return strings.Join([]string{"ShowTopDataRequest", string(data)}, " ")
}
