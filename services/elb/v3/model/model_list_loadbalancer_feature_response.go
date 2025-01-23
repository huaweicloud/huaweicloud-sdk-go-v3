package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListLoadbalancerFeatureResponse Response Object
type ListLoadbalancerFeatureResponse struct {

	// 参数解释：ELB实例特性信息列表。
	Features *[]LoadbalancerFeature `json:"features,omitempty"`

	// 参数解释：请求ID。  注：自动生成 。
	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListLoadbalancerFeatureResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListLoadbalancerFeatureResponse struct{}"
	}

	return strings.Join([]string{"ListLoadbalancerFeatureResponse", string(data)}, " ")
}
