package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteLoadbalancersRequestBody This is a auto create Body Object
type BatchDeleteLoadbalancersRequestBody struct {

	// 解绑后的游离pool是否被删除
	UnboundedPool *bool `json:"unbounded_pool,omitempty"`

	// 是否删除公网IP
	PublicIp *bool `json:"public_ip,omitempty"`

	// 待删除的负载均衡器id列表。
	LoadbalancerIds []string `json:"loadbalancer_ids"`
}

func (o BatchDeleteLoadbalancersRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteLoadbalancersRequestBody struct{}"
	}

	return strings.Join([]string{"BatchDeleteLoadbalancersRequestBody", string(data)}, " ")
}
