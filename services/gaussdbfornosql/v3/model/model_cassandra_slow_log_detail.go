package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CassandraSlowLogDetail struct {

	// 节点名称。
	NodeName string `json:"node_name"`

	// 节点ID。
	NodeId string `json:"node_id"`

	// 执行语句。
	WholeMessage string `json:"whole_message"`

	// 语句类型。
	OperateType string `json:"operate_type"`

	// 执行时间。单位：ms
	CostTime int32 `json:"cost_time"`

	// 数据库键空间
	Keyspace string `json:"keyspace"`

	// 数据库表名
	Table string `json:"table"`

	// 日志产生时间，UTC时间。 格式为“yyyy-mm-ddThh:mm:ssZ”。 其中，T指某个时间的开始；Z指时区偏移量，例如北京时间偏移显示为+0800。
	LogTime string `json:"log_time"`

	// 日志单行序列号
	LineNum string `json:"line_num"`
}

func (o CassandraSlowLogDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CassandraSlowLogDetail struct{}"
	}

	return strings.Join([]string{"CassandraSlowLogDetail", string(data)}, " ")
}
