package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListAuditlogsRequest struct {
	// 语言。

	XLanguage *string `json:"X-Language,omitempty"`
	// 实例ID，可以调用“查询实例列表和详情”接口获取。如果未申请实例，可以调用“创建实例”接口创建。

	InstanceId string `json:"instance_id"`
	// 查询审计日志的节点ID。不传值，默认查询所有的节点,集群实例审计日志分布在mongos节点上。

	NodeId *string `json:"node_id,omitempty"`
	// 查询开始时间，格式为“yyyy-mm-ddThh:mm:ssZ”。其中，T指某个时间的开始；Z指时区偏移量，例如北京时间偏移显示为+0800。

	StartTime string `json:"start_time"`
	// 查询结束时间，格式为“yyyy-mm-ddThh:mm:ssZ”，且大于查询开始时间，时间跨度不超过30天。其中，T指某个时间的开始，Z指时区偏移量，例如北京时间偏移显示为+0800。

	EndTime string `json:"end_time"`
	// 索引位置，偏移量。从第一条数据偏移offset条数据后开始查询，默认为0（偏移0条数据，表示从第一条数据开始查询），必须为数字，不能为负数。

	Offset *int32 `json:"offset,omitempty"`
	// 查询记录数。取值范围：1~100。不传该参数时，默认查询前100条实例信息。

	Limit *int32 `json:"limit,omitempty"`
}

func (o ListAuditlogsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAuditlogsRequest struct{}"
	}

	return strings.Join([]string{"ListAuditlogsRequest", string(data)}, " ")
}
