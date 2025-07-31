package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ParseMultiCloudClusterConfigResponse Response Object
type ParseMultiCloudClusterConfigResponse struct {

	// 集群总数
	TotalNum *int32 `json:"total_num,omitempty"`

	// 多云集群的集群配置列表
	DataList       *[]MultiCloudClusterConfigInfo `json:"data_list,omitempty"`
	HttpStatusCode int                            `json:"-"`
}

func (o ParseMultiCloudClusterConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ParseMultiCloudClusterConfigResponse struct{}"
	}

	return strings.Join([]string{"ParseMultiCloudClusterConfigResponse", string(data)}, " ")
}
