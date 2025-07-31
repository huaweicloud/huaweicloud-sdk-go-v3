package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListMultiCloudClustersResponse Response Object
type ListMultiCloudClustersResponse struct {

	// 集群总数
	TotalNum *int32 `json:"total_num,omitempty"`

	// 多云集群的集群信息列表
	DataList       *[]MultiCloudClusterInfo `json:"data_list,omitempty"`
	HttpStatusCode int                      `json:"-"`
}

func (o ListMultiCloudClustersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListMultiCloudClustersResponse struct{}"
	}

	return strings.Join([]string{"ListMultiCloudClustersResponse", string(data)}, " ")
}
