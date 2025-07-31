package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListClusterAuditLogsResponse Response Object
type ListClusterAuditLogsResponse struct {

	// 日志总条数
	TotalNum *int32 `json:"total_num,omitempty"`

	// k8s集群审计日志列表
	DataList       *[]ClusterAuditLogResponseInfo `json:"data_list,omitempty"`
	HttpStatusCode int                            `json:"-"`
}

func (o ListClusterAuditLogsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListClusterAuditLogsResponse struct{}"
	}

	return strings.Join([]string{"ListClusterAuditLogsResponse", string(data)}, " ")
}
