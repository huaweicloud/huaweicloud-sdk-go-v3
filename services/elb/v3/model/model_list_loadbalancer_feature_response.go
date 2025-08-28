package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListLoadbalancerFeatureResponse Response Object
type ListLoadbalancerFeatureResponse struct {

	// **参数解释**：ELB实例特性信息列表。
	Features *[]LoadbalancerFeature `json:"features,omitempty"`

	// **参数解释**：请求ID。  **取值范围**：由数字、小写字母和中划线（-）组成的字符串，自动生成。
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
