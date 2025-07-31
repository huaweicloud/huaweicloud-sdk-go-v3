package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListContainerLogsResponse Response Object
type ListContainerLogsResponse struct {

	// 日志总条数
	TotalNum *int32 `json:"total_num,omitempty"`

	// k8s集群容器日志列表
	DataList       *[]ClusterContainerLogResponseInfo `json:"data_list,omitempty"`
	HttpStatusCode int                                `json:"-"`
}

func (o ListContainerLogsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListContainerLogsResponse struct{}"
	}

	return strings.Join([]string{"ListContainerLogsResponse", string(data)}, " ")
}
