package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSqlTemplateComparisonsRequestBody 查询SQL模板对比列表请求体
type ListSqlTemplateComparisonsRequestBody struct {

	// 实例ID，实例的唯一标识
	InstanceId string `json:"instance_id"`

	// 对比类型，time 时间段对比，node 节点对比，默认time
	CompareType *string `json:"compare_type,omitempty"`

	// 节点ID，实例节点的唯一标识
	NodeId *string `json:"node_id,omitempty"`

	// 节点对比方1 ID， node模式必选
	NodeId1 *string `json:"node_id1,omitempty"`

	// 节点对比方2 ID， node模式必选
	NodeId2 *string `json:"node_id2,omitempty"`

	// 对比日期1开始时间，单位毫秒
	StartAt1 int64 `json:"start_at1"`

	// 对比日期1结束时间，单位毫秒
	EndAt1 int64 `json:"end_at1"`

	// 对比日期2开始时间，单位毫秒
	StartAt2 int64 `json:"start_at2"`

	// 对比日期2结束时间，单位毫秒
	EndAt2 int64 `json:"end_at2"`

	// 操作类型，可组合，用逗号分隔
	Operation *string `json:"operation,omitempty"`

	// 数据库名称列表
	DbNameList *[]string `json:"db_name_list,omitempty"`

	// 关键字，模糊搜索
	Keyword *string `json:"keyword,omitempty"`

	// SQL模板ID
	SqlTemplateId *string `json:"sql_template_id,omitempty"`

	// 排序字段，取值范围：executeNum（执行次数）、totalCost（总耗时）、avgCost（平均耗时）、totalScan（总扫描行数）、avgScan（平均扫描行数）
	Sort *string `json:"sort,omitempty"`

	// 排序顺序，true（正序）、false（倒序）
	Asc *bool `json:"asc,omitempty"`

	// 单次查询数量
	Size *int32 `json:"size,omitempty"`
}

func (o ListSqlTemplateComparisonsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSqlTemplateComparisonsRequestBody struct{}"
	}

	return strings.Join([]string{"ListSqlTemplateComparisonsRequestBody", string(data)}, " ")
}
