package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 慢日志请求体。
type LtsLogSlowQueryRequest struct {

	// 节点ID。
	NodeId string `json:"node_id"`

	// 开始时间，格式为“yyyy-mm-ddThh:mm:ssZ”。  其中，T指某个时间的开始；Z指时区偏移量，例如偏移1个小时显示为+0100。
	StartTime string `json:"start_time"`

	// 结束时间，格式为“yyyy-mm-ddThh:mm:ssZ”。  其中，T指某个时间的开始；Z指时区偏移量，例如偏移1个小时显示为+0100。
	EndTime string `json:"end_time"`

	// 查询记录数。
	Limit int32 `json:"limit"`

	// 日志单行序列号，第一次查询时不需要此参数,后续分页查询时需要使用,可从上次查询的返回信息中获取。
	LineNum *string `json:"line_num,omitempty"`

	// SQL语句类型，取空值，表示查询所有语句类型，也可指定如下日志类型： - INSERT - UPDATE - SELECT - DELETE - ALTER - DROP - CREATE
	OperateType *string `json:"operate_type,omitempty"`

	// 所属数据库。
	Database *string `json:"database,omitempty"`
}

func (o LtsLogSlowQueryRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LtsLogSlowQueryRequest struct{}"
	}

	return strings.Join([]string{"LtsLogSlowQueryRequest", string(data)}, " ")
}
