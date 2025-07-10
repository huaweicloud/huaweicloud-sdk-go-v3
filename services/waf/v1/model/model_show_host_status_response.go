package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowHostStatusResponse Response Object
type ShowHostStatusResponse struct {

	// **参数解释：** 域名 **取值范围：** 不涉及
	Name *string `json:"name,omitempty"`

	// **参数解释：** 域名的防护状态 **取值范围：** - enabled：防护中，WAF根据您配置的策略进行攻击检测 - disabled：未防护，WAF只转发该域名的请求，不做攻击检测 - bypassed：该域名的请求直接到达其后端服务器，不再经过WAF
	Status *string `json:"status,omitempty"`

	// **参数解释：** 域名ID **取值范围：** 不涉及
	WafInstanceId  *string `json:"waf_instance_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowHostStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowHostStatusResponse struct{}"
	}

	return strings.Join([]string{"ShowHostStatusResponse", string(data)}, " ")
}
