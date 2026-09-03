package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTemplateDatabaseComparisonsRequest Request Object
type ListTemplateDatabaseComparisonsRequest struct {

	// 实例ID，实例的唯一标识
	InstanceId string `json:"instance_id"`

	// 节点ID，实例节点的唯一标识
	NodeId *string `json:"node_id,omitempty"`

	// 比较时间段1的开始时间，Unix timestamp，单位：毫秒
	StartAt1 int64 `json:"start_at1"`

	// 比较时间段1的结束时间，Unix timestamp，单位：毫秒
	EndAt1 int64 `json:"end_at1"`

	// 比较时间段2的开始时间，Unix timestamp，单位：毫秒
	StartAt2 int64 `json:"start_at2"`

	// 比较时间段2的结束时间，Unix timestamp，单位：毫秒
	EndAt2 int64 `json:"end_at2"`

	// 操作类型，可组合，用逗号分隔
	Operation *string `json:"operation,omitempty"`

	// 数据库列表
	DbNameList *[]string `json:"db_name_list,omitempty"`

	// 关键字
	Keyword *string `json:"keyword,omitempty"`

	// 排序字段，取值范围：executeNum（执行次数）、totalCost（总耗时）、avgCost（平均耗时）、totalScan（总扫描行数）、avgScan（平均扫描行数）
	Sort *string `json:"sort,omitempty"`

	// 排序顺序，true（正序）、false（逆序）
	Asc *bool `json:"asc,omitempty"`

	// 数量，默认30
	Size *int32 `json:"size,omitempty"`
}

func (o ListTemplateDatabaseComparisonsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTemplateDatabaseComparisonsRequest struct{}"
	}

	return strings.Join([]string{"ListTemplateDatabaseComparisonsRequest", string(data)}, " ")
}
