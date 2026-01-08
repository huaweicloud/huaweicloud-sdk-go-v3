package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AzAffinity struct {

	// **参数解释**：后端服务器组是否开启可用区亲和。开启后，负载均衡器会按照配置的亲和策略进行转发。  **约束限制**： - 仅IP,TCP,UDP的pool支持。 - TLS监听器绑定的pool不支持。 - 开启后，原本的pool_health配置失效。  **取值范围**：false不开启，true开启。  **默认取值**：不涉及
	Enable bool `json:"enable"`

	// **参数解释**：后端服务器组单可用区百分比健康度最小阈值，当“后端服务器组单可用区百分比健康度”小于该阈值时，触发可用区异常退避策略。“后端服务器组单可用区百分比健康度”是指在一个后端服务器组中，同可用区中健康检查结果正常的服务器数量与该后端服务器组中属于该可用区的后端服务器总数量的比值，百分比结果向上取整。例如：后端服务器组中属于可用区A的后端服务器总数量为3，设置后端服务器组单可用区百分比健康度最小阈值为66时，3x0.66=1.98向上取整为数量阈值2台，即属于可用区A的健康后端数小于2台时触发退避策略；设置后端服务器组单可用区百分比健康度最小阈值为67时，3x0.67=2.01向上取整为数量阈值3台，即属于可用区A的健康后端数小于3台时触发退避策略。  **约束限制**： - enable为true时，az_minimum_healthy_member_percentage与az_minimum_healthy_member_count不能同时为-1 - enable为true时，az_minimum_healthy_member_percentage与az_minimum_healthy_member_count不能同时不为-1  **取值范围**：-1至100的整数，0-100为百分比范围，-1表示采用数量阈值。  **默认取值**：不涉及
	AzMinimumHealthyMemberPercentage *int32 `json:"az_minimum_healthy_member_percentage,omitempty"`

	// **参数解释**：后端服务器组单可用区中数量健康度最小阈值，当“后端服务器组单可用区中数量健康度”小于该阈值时，触发可用区异常退避策略。“后端服务器组单可用区中数量健康度”是指在一个后端服务器组中，属于同一个可用区的健康检查结果正常的服务器数量。  **约束限制**： - enable为true时，az_minimum_healthy_member_percentage与az_minimum_healthy_member_count不能同时为-1 - enable为true时，az_minimum_healthy_member_percentage与az_minimum_healthy_member_count不能同时不为-1  **取值范围**：-1至10000的整数，0-10000为数量范围，-1表示采用百分比阈值。  **默认取值**：不涉及
	AzMinimumHealthyMemberCount *int32 `json:"az_minimum_healthy_member_count,omitempty"`

	// **参数解释**：后端服务器组单可用区异常退避策略。后端服务器组的健康度小于所配置的最小阈值时，触发该退避策略。forward_to_all_member_of_local_az：转发至同可用区的所有后端服务器，无论健康检查结果是否正常；forward_to_healthy_member_of_remote_az：转发至非同可用区中所有健康检查结果正常的后端服务器；forward_to_all_healthy_member：转发至所有可用区中健康检查结果正常的后端服务器；forward_to_all_member：转发至所有可用区中的所有后端服务器，无论健康检查结果是否正常  **约束限制**：不涉及  **取值范围**：forward_to_all_member_of_local_az，forward_to_healthy_member_of_remote_az，forward_to_all_healthy_member，forward_to_all_member。  **默认取值**：forward_to_all_member_of_local_az
	AzUnhealthyFallbackStrategy *string `json:"az_unhealthy_fallback_strategy,omitempty"`
}

func (o AzAffinity) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AzAffinity struct{}"
	}

	return strings.Join([]string{"AzAffinity", string(data)}, " ")
}
