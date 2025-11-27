package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListClusterProtectionInfoResponse Response Object
type ListClusterProtectionInfoResponse struct {

	// **参数解释**: 总数 **取值范围**: 最小值0，最大值10000
	TotalNum *int32 `json:"total_num,omitempty"`

	// **参数解释**: 集群防护信息列表 **取值范围**: 取值0-10000个ClusterResponseInfo对象
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
