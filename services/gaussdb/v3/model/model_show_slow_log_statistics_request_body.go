package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ShowSlowLogStatisticsRequestBody struct {

	// 每页多少条记录（查询结果），取值范围是1~100，不填时默认为10。
	Limit *int32 `json:"limit,omitempty"`

	// 索引位置，偏移量。默认为0，表示从第一条数据开始查询。
	Offset *int32 `json:"offset,omitempty"`

	// 节点ID。
	NodeId string `json:"node_id"`

	// 语句类型，取空值，表示查询所有语句类型。  枚举值:   - INSERT   - UPDATE   - SELECT   - DELETE   - CREATE   - ALL
	Type *string `json:"type,omitempty"`

	// 数据库名称。数据库名称不支持包含特殊字符 < > & 等的搜索。
	Database *string `json:"database,omitempty"`

	// 开始日期，格式为“yyyy-mm-ddThh:mm:ssZ”。  其中，T指某个时间的开始；Z指时区偏移量，例如北京时间偏移显示为+0800。
	StartTime string `json:"start_time"`

	// 结束时间，格式为“yyyy-mm-ddThh:mm:ssZ”。  其中，T指某个时间的开始；Z指时区偏移量，例如北京时间偏移显示为+0800。只能查询当前时间前一个月内的慢日志。
	EndTime string `json:"end_time"`

	// 指定排序字段。   - executeTime：表示按照平均执行时间降序排序。   - 字段为空或传入其他值，表示按照执行次数降序排序。
	Sort *string `json:"sort,omitempty"`

	// 排序顺序。默认desc。 枚举值：   - desc   - asc
	Order *string `json:"order,omitempty"`
}

func (o ShowSlowLogStatisticsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSlowLogStatisticsRequestBody struct{}"
	}

	return strings.Join([]string{"ShowSlowLogStatisticsRequestBody", string(data)}, " ")
}
