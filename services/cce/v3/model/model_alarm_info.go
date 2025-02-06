package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AlarmInfo **参数解释：** 告警助手参数配置。基于AOM服务的告警能力实现，提供集群内的告警快速检索、告警快速配置的能力，告警中心的指标类告警规则依赖云原生监控插件上报数据到AOM实例。 **约束限制：** 不涉及
type AlarmInfo struct {

	// **参数解释：** 联系组列表。填写SMN主题名称，通过配置告警联系组，分组管理订阅终端，接收告警信息。 **约束限制：** 不涉及
	Topics []string `json:"topics"`

	// **参数解释：** 开启告警助手时传入告警模板ID。默认采用容器场景下的告警规则模板。 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	AlarmRuleTemplateId *string `json:"alarmRuleTemplateId,omitempty"`

	// **参数解释：** 开启告警助手时传入AOM普罗实例的id。若未安装普罗插件或者未对接AOM实例，此参数无需指定，告警中心将不会创建指标类告警规则。 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	PromInstanceID *string `json:"promInstanceID,omitempty"`

	// **参数解释：** 开启告警助手时传入AOM普罗实例的企业项目id。若未安装普罗插件或者未对接AOM实例，此参数无需指定，告警中心将不会创建指标类告警规则。 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	PromEnterpriseProjectID *string `json:"promEnterpriseProjectID,omitempty"`
}

func (o AlarmInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AlarmInfo struct{}"
	}

	return strings.Join([]string{"AlarmInfo", string(data)}, " ")
}
