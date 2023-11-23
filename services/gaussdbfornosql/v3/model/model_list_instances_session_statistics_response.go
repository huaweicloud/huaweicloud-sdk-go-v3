package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListInstancesSessionStatisticsResponse Response Object
type ListInstancesSessionStatisticsResponse struct {

	// 总客户端连接数。
	TotalConnectionCount *int32 `json:"total_connection_count,omitempty"`

	// active_connection_count
	ActiveConnectionCount *int32 `json:"active_connection_count,omitempty"`

	// 节点连接的各个客户端连接数汇总，从大到小取前十个，最多十个，展示客户端的ip地址和连接总数。
	TopSourceIps *[]ListInstancesSessionStatisticsRespondBodyTopSourceIps `json:"top_source_ips,omitempty"`

	// 节点各数据库连接的客户端的ip和该ip连接节点的连接数，按连接数从高到低取前十个，最多十个。
	TopDbs         *interface{} `json:"top_dbs,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ListInstancesSessionStatisticsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstancesSessionStatisticsResponse struct{}"
	}

	return strings.Join([]string{"ListInstancesSessionStatisticsResponse", string(data)}, " ")
}
