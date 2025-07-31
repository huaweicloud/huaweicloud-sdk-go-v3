package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListClusterProtectionInfoResponse Response Object
type ListClusterProtectionInfoResponse struct {

	// 总数
	TotalNum *int32 `json:"total_num,omitempty"`

	// 集群防护信息列表
	DataList       *[]ClusterResponseInfo `json:"data_list,omitempty"`
	HttpStatusCode int                    `json:"-"`
}

func (o ListClusterProtectionInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListClusterProtectionInfoResponse struct{}"
	}

	return strings.Join([]string{"ListClusterProtectionInfoResponse", string(data)}, " ")
}
