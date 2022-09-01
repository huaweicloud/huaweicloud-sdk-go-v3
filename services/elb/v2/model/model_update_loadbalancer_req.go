package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 更新负载均衡器的请求体
type UpdateLoadbalancerReq struct {

	// 负载均衡器名称。
	Name *string `json:"name,omitempty" xml:"name"`

	// 负载均衡器的描述信息
	Description *string `json:"description,omitempty" xml:"description"`

	// 负载均衡器的管理状态。只支持设定为true，该字段的值无实际意义。
	AdminStateUp *bool `json:"admin_state_up,omitempty" xml:"admin_state_up"`
}

func (o UpdateLoadbalancerReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateLoadbalancerReq struct{}"
	}

	return strings.Join([]string{"UpdateLoadbalancerReq", string(data)}, " ")
}
