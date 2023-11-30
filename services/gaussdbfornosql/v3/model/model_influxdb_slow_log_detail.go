package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type InfluxdbSlowLogDetail struct {

	// 节点ID。
	NodeId string `json:"node_id"`

	// 节点名称。
	NodeName string `json:"node_name"`

	// 执行语句。
	WholeMessage string `json:"whole_message"`

	// 语句类型。
	OperateType string `json:"operate_type"`

	// 执行时间。单位：ms
	CostTime string `json:"cost_time"`

	// 日志产生时间，UTC时间。 格式为“yyyy-mm-ddThh:mm:ssZ”。 其中，T指某个时间的开始；Z指时区偏移量，例如北京时间偏移显示为+0800。
	LogTime string `json:"log_time"`

	// 数据库名称。
	Database string `json:"database"`

	// 保留策略。
	RetentionPolicy string `json:"retention_policy"`

	// 日志单行序列号。
	LineNum string `json:"line_num"`
}

func (o InfluxdbSlowLogDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InfluxdbSlowLogDetail struct{}"
	}

	return strings.Join([]string{"InfluxdbSlowLogDetail", string(data)}, " ")
}
