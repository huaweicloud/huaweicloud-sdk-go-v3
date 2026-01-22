package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateFirewallNameDto struct {

	// **参数解释**： 防火墙名称 **约束限制**： 防火墙名称为中英文，包含数字，下划线，中划线，长度为1-48 **取值范围**： 不涉及 **默认取值**： 不涉及
	FwInstanceName string `json:"fw_instance_name"`
}

func (o UpdateFirewallNameDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateFirewallNameDto struct{}"
	}

	return strings.Join([]string{"UpdateFirewallNameDto", string(data)}, " ")
}
