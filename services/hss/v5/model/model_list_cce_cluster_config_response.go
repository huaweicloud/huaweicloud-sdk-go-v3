package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCceClusterConfigResponse Response Object
type ListCceClusterConfigResponse struct {

	// 配置总数
	TotalNum *int32 `json:"total_num,omitempty"`

	// 配置信息列表
	DataList       *[]ClusterConfigResponseInfo `json:"data_list,omitempty"`
	HttpStatusCode int                          `json:"-"`
}

func (o ListCceClusterConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCceClusterConfigResponse struct{}"
	}

	return strings.Join([]string{"ListCceClusterConfigResponse", string(data)}, " ")
}
