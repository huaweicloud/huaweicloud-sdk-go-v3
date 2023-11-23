package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ErrorLogDetail struct {

	// 节点名称。
	NodeName string `json:"node_name"`

	// 节点ID。
	NodeId string `json:"node_id"`

	// 描述信息。
	RawMessage string `json:"raw_message"`

	// 日志级别。
	Severity string `json:"severity"`

	// 日志产生时间，UTC时间。 格式为“yyyy-mm-ddThh:mm:ssZ”。 其中，T指某个时间的开始；Z指时区偏移量，例如北京时间偏移显示为+0800。
	LogTime string `json:"log_time"`

	// 日志单行序列号。
	LineNum string `json:"line_num"`
}

func (o ErrorLogDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ErrorLogDetail struct{}"
	}

	return strings.Join([]string{"ErrorLogDetail", string(data)}, " ")
}
