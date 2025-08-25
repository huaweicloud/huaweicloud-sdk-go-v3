package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowCcspClusterListResponse Response Object
type ShowCcspClusterListResponse struct {

	// 满足条件的集群总数
	TotalNum *int32 `json:"total_num,omitempty"`

	// 集群列表
	Result         *[]CcspClusterInfo `json:"result,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o ShowCcspClusterListResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCcspClusterListResponse struct{}"
	}

	return strings.Join([]string{"ShowCcspClusterListResponse", string(data)}, " ")
}
