package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListContainerNetworkClustersResponse Response Object
type ListContainerNetworkClustersResponse struct {

	// 总数
	TotalNum *int32 `json:"total_num,omitempty"`

	// 集群信息列表
	DataList       *[]ClustersResponseInfo `json:"data_list,omitempty"`
	HttpStatusCode int                     `json:"-"`
}

func (o ListContainerNetworkClustersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListContainerNetworkClustersResponse struct{}"
	}

	return strings.Join([]string{"ListContainerNetworkClustersResponse", string(data)}, " ")
}
