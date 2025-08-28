package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RedirectPoolsConfig 转发策略服务器组的权重配置。当监听器的高级转发策略功能（enhance_l7policy_enable）开启后才会生效，未开启传入该字段会报错。  [共享型负载均衡器下的转发策略不支持该字段，传入会报错。](tag:hws,hws_hk,ocb,ctc,hcs,g42,tm,cmcc,hk_g42,hws_ocb,srg,fcs,dt)  [不支持该字段，请勿使用。](tag:hcso_dt)  [荷兰region不支持该字段，请勿使用。](tag:dt)
type RedirectPoolsConfig struct {

	// **参数解释**：所在后端服务器组ID。  **取值范围**：不涉及
	PoolId *string `json:"pool_id,omitempty"`

	// **参数解释**：转发策略服务器组的权重。请求将根据该权重进行负载分发到不同的服务器组。权重值越大，分发的请求越多。权重为0的后端不再接受新的请求。  **取值范围**：0-100
	Weight *int32 `json:"weight,omitempty"`
}

func (o RedirectPoolsConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RedirectPoolsConfig struct{}"
	}

	return strings.Join([]string{"RedirectPoolsConfig", string(data)}, " ")
}
