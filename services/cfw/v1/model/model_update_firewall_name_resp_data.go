package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateFirewallNameRespData **参数解释**： 响应信息 **取值范围**： 不涉及
type UpdateFirewallNameRespData struct {

	// **参数解释**： 防火墙ID，可通过防火墙ID获取 方式获取 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	Id *string `json:"id,omitempty"`

	// **参数解释**： 防火墙名称 **约束限制**： 防火墙名称为中英文，包含数字，下划线，中划线，长度为1-48 **取值范围**： 不涉及 **默认取值**： 不涉及
	Name *string `json:"name,omitempty"`
}

func (o UpdateFirewallNameRespData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateFirewallNameRespData struct{}"
	}

	return strings.Join([]string{"UpdateFirewallNameRespData", string(data)}, " ")
}
