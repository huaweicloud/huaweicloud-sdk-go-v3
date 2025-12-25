package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AutoGenerateSecurityGroupHardeningConfigSpec struct {

	// **参数解释：** 自动创建安全组时可选择不开放node节点相关端口，支持单端口和端口段两种形式。示例如下： ``` \"portsToDisable\": [\"22\", \"4090-4099\"] ``` **约束限制：** 若指定了节点安全组SecurityGroup，该配置项将被忽略。 只针对CCE Standard和Turbo集群的node安全组生效，不支持master安全组，不支持eni安全组。 **取值范围：** 端口号仅支持整数，范围[1-65535] **默认取值：** 不涉及
	PortsToDisable *[]string `json:"portsToDisable,omitempty"`
}

func (o AutoGenerateSecurityGroupHardeningConfigSpec) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AutoGenerateSecurityGroupHardeningConfigSpec struct{}"
	}

	return strings.Join([]string{"AutoGenerateSecurityGroupHardeningConfigSpec", string(data)}, " ")
}
