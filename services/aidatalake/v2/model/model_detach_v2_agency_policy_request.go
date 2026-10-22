package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// DetachV2AgencyPolicyRequest Request Object
type DetachV2AgencyPolicyRequest struct {

	// **参数解释**：委托名称。 **约束限制**：固定值ai_datalake_admin_trust。 **取值范围**：不涉及。 **默认取值**：不涉及。
	AgencyName string `json:"agency_name"`

	// **参数解释**：权限策略名称，使用AI DataLake的功能，需要授权一些权限策略。 **约束限制**：不支持移除必选策略。 **取值范围**：   - AI_DATALAKE_SMN_POLICY：消息通知功能相关权限策略，用来将系统通知消息转发到SMN。   - AI_DATALAKE_AOM_POLICY：AOM功能相关权限策略。   - AI_DATALAKE_SWR_POLICY：镜像包注册功能相关权限策略。   - AI_DATALAKE_VPCEP_POLICY：网络打通功能相关权限策略。   - AI_DATALAKE_DNS_POLICY：DNS相关权限策略。   - AI_DATALAKE_OBS_POLICY：OBS相关权限策略。   - AI_DATALAKE_VPC_POLICY：VPC相关权限策略。   - AI_DATALAKE_DCS_POLICY：DCS相关权限策略。   - AI_DATALAKE_DEW_POLICY：DEW相关权限策略。 **默认取值**：不涉及。
	PolicyName DetachV2AgencyPolicyRequestPolicyName `json:"policy_name"`
}

func (o DetachV2AgencyPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DetachV2AgencyPolicyRequest struct{}"
	}

	return strings.Join([]string{"DetachV2AgencyPolicyRequest", string(data)}, " ")
}

type DetachV2AgencyPolicyRequestPolicyName struct {
	value string
}

type DetachV2AgencyPolicyRequestPolicyNameEnum struct {
	AI_DATALAKE_SMN_POLICY   DetachV2AgencyPolicyRequestPolicyName
	AI_DATALAKE_AOM_POLICY   DetachV2AgencyPolicyRequestPolicyName
	AI_DATALAKE_SWR_POLICY   DetachV2AgencyPolicyRequestPolicyName
	AI_DATALAKE_VPCEP_POLICY DetachV2AgencyPolicyRequestPolicyName
	AI_DATALAKE_DNS_POLICY   DetachV2AgencyPolicyRequestPolicyName
	AI_DATALAKE_OBS_POLICY   DetachV2AgencyPolicyRequestPolicyName
	AI_DATALAKE_VPC_POLICY   DetachV2AgencyPolicyRequestPolicyName
	AI_DATALAKE_DCS_POLICY   DetachV2AgencyPolicyRequestPolicyName
	AI_DATALAKE_DEW_POLICY   DetachV2AgencyPolicyRequestPolicyName
}

func GetDetachV2AgencyPolicyRequestPolicyNameEnum() DetachV2AgencyPolicyRequestPolicyNameEnum {
	return DetachV2AgencyPolicyRequestPolicyNameEnum{
		AI_DATALAKE_SMN_POLICY: DetachV2AgencyPolicyRequestPolicyName{
			value: "AI_DATALAKE_SMN_POLICY",
		},
		AI_DATALAKE_AOM_POLICY: DetachV2AgencyPolicyRequestPolicyName{
			value: "AI_DATALAKE_AOM_POLICY",
		},
		AI_DATALAKE_SWR_POLICY: DetachV2AgencyPolicyRequestPolicyName{
			value: "AI_DATALAKE_SWR_POLICY",
		},
		AI_DATALAKE_VPCEP_POLICY: DetachV2AgencyPolicyRequestPolicyName{
			value: "AI_DATALAKE_VPCEP_POLICY",
		},
		AI_DATALAKE_DNS_POLICY: DetachV2AgencyPolicyRequestPolicyName{
			value: "AI_DATALAKE_DNS_POLICY",
		},
		AI_DATALAKE_OBS_POLICY: DetachV2AgencyPolicyRequestPolicyName{
			value: "AI_DATALAKE_OBS_POLICY",
		},
		AI_DATALAKE_VPC_POLICY: DetachV2AgencyPolicyRequestPolicyName{
			value: "AI_DATALAKE_VPC_POLICY",
		},
		AI_DATALAKE_DCS_POLICY: DetachV2AgencyPolicyRequestPolicyName{
			value: "AI_DATALAKE_DCS_POLICY",
		},
		AI_DATALAKE_DEW_POLICY: DetachV2AgencyPolicyRequestPolicyName{
			value: "AI_DATALAKE_DEW_POLICY",
		},
	}
}

func (c DetachV2AgencyPolicyRequestPolicyName) Value() string {
	return c.value
}

func (c DetachV2AgencyPolicyRequestPolicyName) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DetachV2AgencyPolicyRequestPolicyName) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}
