package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

// V2AgencyPolicy **参数解释**：委托中授权的权限策略。 **取值范围**：不涉及。
type V2AgencyPolicy struct {

	// **参数解释**：权限策略名称，指定授权给AI DataLake服务的权限策略。 **取值范围**：   - AI_DATALAKE_COMMON_POLICY：基础服务相关权限策略。   - AI_DATALAKE_SELF_POLICY：服务自身访问权限策略，用于服务内部资源访问。   - AI_DATALAKE_SMN_POLICY：消息通知功能相关权限策略，用来将系统通知消息转发到SMN。   - AI_DATALAKE_LAKEFORMATION_POLICY：LakeFormation功能相关权限策略。   - AI_DATALAKE_AOM_POLICY：AOM功能相关权限策略。   - AI_DATALAKE_SWR_POLICY：镜像包注册功能相关权限策略。   - AI_DATALAKE_LTS_POLICY：日志功能相关权限策略。   - AI_DATALAKE_VPCEP_POLICY：网络打通功能相关权限策略。   - AI_DATALAKE_DNS_POLICY：DNS相关权限策略。   - AI_DATALAKE_OBS_POLICY：OBS相关权限策略。   - AI_DATALAKE_VPC_POLICY：VPC相关权限策略。   - AI_DATALAKE_DCS_POLICY：DCS相关权限策略。   - AI_DATALAKE_DEW_POLICY：DEW相关权限策略。
	PolicyName *V2AgencyPolicyPolicyName `json:"policy_name,omitempty"`

	// **参数解释**：权限策略名称。 **取值范围**：支持中文字符、英文字母、数字及特殊字符，长度0~128个字符。
	DisplayName *string `json:"display_name,omitempty"`

	// **参数解释**：是否是必需的。 **取值范围**：   - true：必需。   - false：非必需。
	IsRequired bool `json:"is_required"`

	// **参数解释**：是否已开启。只要授权过（策略存在且 effect 为 Allow），即为 true。 **取值范围**：   - true：已开启。   - false：未开启。
	IsEnable bool `json:"is_enable"`

	// **参数解释**：策略描述。 **取值范围**：支持中文字符、英文字母、数字及特殊字符，长度0~128个字符。
	Description string `json:"description"`

	// **参数解释**：权限策略的action详情列表，详细说明该策略包含的权限操作。
	ActionDetails []ActionDetail `json:"action_details"`

	// **参数解释**：资源创建时间。 **取值范围**：ISO 8601格式，yyyy-MM-dd'T'HH:mm:ss.SSSZ，UTC时间。
	CreateTime *sdktime.SdkTime `json:"create_time"`

	// **参数解释**：策略内容与预期不一致时的 action 变更详情列表。只展示需要新增的 action（代码有但 IAM 没有的）。策略内容完全一致或策略完全缺失时返回空数组。
	ChangedActionDetails []ChangedActionDetail `json:"changed_action_details"`
}

func (o V2AgencyPolicy) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "V2AgencyPolicy struct{}"
	}

	return strings.Join([]string{"V2AgencyPolicy", string(data)}, " ")
}

type V2AgencyPolicyPolicyName struct {
	value string
}

type V2AgencyPolicyPolicyNameEnum struct {
	AI_DATALAKE_COMMON_POLICY        V2AgencyPolicyPolicyName
	AI_DATALAKE_SELF_POLICY          V2AgencyPolicyPolicyName
	AI_DATALAKE_SMN_POLICY           V2AgencyPolicyPolicyName
	AI_DATALAKE_LAKEFORMATION_POLICY V2AgencyPolicyPolicyName
	AI_DATALAKE_AOM_POLICY           V2AgencyPolicyPolicyName
	AI_DATALAKE_SWR_POLICY           V2AgencyPolicyPolicyName
	AI_DATALAKE_LTS_POLICY           V2AgencyPolicyPolicyName
	AI_DATALAKE_VPCEP_POLICY         V2AgencyPolicyPolicyName
	AI_DATALAKE_DNS_POLICY           V2AgencyPolicyPolicyName
	AI_DATALAKE_OBS_POLICY           V2AgencyPolicyPolicyName
	AI_DATALAKE_VPC_POLICY           V2AgencyPolicyPolicyName
	AI_DATALAKE_DCS_POLICY           V2AgencyPolicyPolicyName
	AI_DATALAKE_DEW_POLICY           V2AgencyPolicyPolicyName
}

func GetV2AgencyPolicyPolicyNameEnum() V2AgencyPolicyPolicyNameEnum {
	return V2AgencyPolicyPolicyNameEnum{
		AI_DATALAKE_COMMON_POLICY: V2AgencyPolicyPolicyName{
			value: "AI_DATALAKE_COMMON_POLICY",
		},
		AI_DATALAKE_SELF_POLICY: V2AgencyPolicyPolicyName{
			value: "AI_DATALAKE_SELF_POLICY",
		},
		AI_DATALAKE_SMN_POLICY: V2AgencyPolicyPolicyName{
			value: "AI_DATALAKE_SMN_POLICY",
		},
		AI_DATALAKE_LAKEFORMATION_POLICY: V2AgencyPolicyPolicyName{
			value: "AI_DATALAKE_LAKEFORMATION_POLICY",
		},
		AI_DATALAKE_AOM_POLICY: V2AgencyPolicyPolicyName{
			value: "AI_DATALAKE_AOM_POLICY",
		},
		AI_DATALAKE_SWR_POLICY: V2AgencyPolicyPolicyName{
			value: "AI_DATALAKE_SWR_POLICY",
		},
		AI_DATALAKE_LTS_POLICY: V2AgencyPolicyPolicyName{
			value: "AI_DATALAKE_LTS_POLICY",
		},
		AI_DATALAKE_VPCEP_POLICY: V2AgencyPolicyPolicyName{
			value: "AI_DATALAKE_VPCEP_POLICY",
		},
		AI_DATALAKE_DNS_POLICY: V2AgencyPolicyPolicyName{
			value: "AI_DATALAKE_DNS_POLICY",
		},
		AI_DATALAKE_OBS_POLICY: V2AgencyPolicyPolicyName{
			value: "AI_DATALAKE_OBS_POLICY",
		},
		AI_DATALAKE_VPC_POLICY: V2AgencyPolicyPolicyName{
			value: "AI_DATALAKE_VPC_POLICY",
		},
		AI_DATALAKE_DCS_POLICY: V2AgencyPolicyPolicyName{
			value: "AI_DATALAKE_DCS_POLICY",
		},
		AI_DATALAKE_DEW_POLICY: V2AgencyPolicyPolicyName{
			value: "AI_DATALAKE_DEW_POLICY",
		},
	}
}

func (c V2AgencyPolicyPolicyName) Value() string {
	return c.value
}

func (c V2AgencyPolicyPolicyName) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *V2AgencyPolicyPolicyName) UnmarshalJSON(b []byte) error {
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
