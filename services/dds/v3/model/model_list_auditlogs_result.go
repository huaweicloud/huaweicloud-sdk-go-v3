package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListAuditlogsResult struct {

	// 节点ID。
	NodeId string `json:"node_id"`

	// 审计日志ID。
	Id string `json:"id"`

	// 审计日志文件名。
	Name string `json:"name"`

	// 审计日志大小，单位：byte。
	Size int64 `json:"size"`

	// 审计日志开始时间，格式为“yyyy-mm-ddThh:mm:ssZ”。其中，T指某个时间的开始，Z指时区偏移量，例如北京时间偏移显示为+0800。
	StartTime string `json:"start_time"`

	// 审计日志结束时间，格式为“yyyy-mm-ddThh:mm:ssZ”。其中，T指某个时间的开始；Z指时区偏移量，例如北京时间偏移显示为+0800。
	EndTime string `json:"end_time"`
}

func (o ListAuditlogsResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAuditlogsResult struct{}"
	}

	return strings.Join([]string{"ListAuditlogsResult", string(data)}, " ")
}
