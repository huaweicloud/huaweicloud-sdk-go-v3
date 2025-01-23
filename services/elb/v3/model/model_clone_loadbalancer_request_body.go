package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CloneLoadbalancerRequestBody 复制负载均衡器接口请求体。
type CloneLoadbalancerRequestBody struct {

	// 参数解释：单次最大复制数量。  取值范围：1-10  默认取值：1
	Count *int32 `json:"count,omitempty"`

	TargetLoadbalancerParam *TargetLoadbalancerParam `json:"target_loadbalancer_param"`
}

func (o CloneLoadbalancerRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CloneLoadbalancerRequestBody struct{}"
	}

	return strings.Join([]string{"CloneLoadbalancerRequestBody", string(data)}, " ")
}
