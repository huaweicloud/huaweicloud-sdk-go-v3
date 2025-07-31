package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListClusterEventLogsResponse Response Object
type ListClusterEventLogsResponse struct {

	// 日志总条数
	TotalNum *int32 `json:"total_num,omitempty"`

	// k8s集群事件日志列表
	DataList       *[]ClusterEventLogResponseInfo `json:"data_list,omitempty"`
	HttpStatusCode int                            `json:"-"`
}

func (o ListClusterEventLogsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListClusterEventLogsResponse struct{}"
	}

	return strings.Join([]string{"ListClusterEventLogsResponse", string(data)}, " ")
}
