package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateV2AgencyRequestBody **参数解释**：创建委托请求参数。 **约束限制**：必须包含以下必选策略：   - AI_DATALAKE_COMMON_POLICY   - AI_DATALAKE_SELF_POLICY   - AI_DATALAKE_LAKEFORMATION_POLICY   - AI_DATALAKE_LTS_POLICY **取值范围**：不涉及。 **默认取值**：不涉及。
type CreateV2AgencyRequestBody struct {

	// **参数解释**：权限策略列表，指定需要授权给AI DataLake服务的权限策略，取值包括：   - AI_DATALAKE_COMMON_POLICY：基础服务权限策略。   - AI_DATALAKE_SELF_POLICY：AIDataLake服务自身访问权限策略，用于服务内部资源访问。   - AI_DATALAKE_LAKEFORMATION_POLICY：LakeFormation功能相关权限策略。   - AI_DATALAKE_LTS_POLICY：日志功能相关权限策略。   - AI_DATALAKE_SMN_POLICY：消息通知功能相关权限策略，用来将系统通知消息转发到SMN。   - AI_DATALAKE_AOM_POLICY：AOM功能相关权限策略。   - AI_DATALAKE_SWR_POLICY：镜像包注册功能相关权限策略。   - AI_DATALAKE_VPCEP_POLICY：网络打通功能相关权限策略。   - AI_DATALAKE_DNS_POLICY：DNS相关权限策略。   - AI_DATALAKE_OBS_POLICY：OBS相关权限策略。   - AI_DATALAKE_VPC_POLICY：VPC相关权限策略。   - AI_DATALAKE_DCS_POLICY：DCS相关权限策略。   - AI_DATALAKE_DEW_POLICY：DEW相关权限策略。 **约束限制**：创建委托时，必须包含以下四个必需策略：   - AI_DATALAKE_COMMON_POLICY：基础服务权限策略。   - AI_DATALAKE_SELF_POLICY：AIDataLake服务自身访问权限策略，用于服务内部资源访问。   - AI_DATALAKE_LAKEFORMATION_POLICY：LakeFormation功能相关权限策略。   - AI_DATALAKE_LTS_POLICY：日志功能相关权限策略。
	PolicyNames []CreateV2AgencyRequestBodyPolicyNames `json:"policy_names"`
}

func (o CreateV2AgencyRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateV2AgencyRequestBody struct{}"
	}

	return strings.Join([]string{"CreateV2AgencyRequestBody", string(data)}, " ")
}

type CreateV2AgencyRequestBodyPolicyNames struct {
	value string
}

type CreateV2AgencyRequestBodyPolicyNamesEnum struct {
	AI_DATALAKE_COMMON_POLICY        CreateV2AgencyRequestBodyPolicyNames
	AI_DATALAKE_SELF_POLICY          CreateV2AgencyRequestBodyPolicyNames
	AI_DATALAKE_SMN_POLICY           CreateV2AgencyRequestBodyPolicyNames
	AI_DATALAKE_LAKEFORMATION_POLICY CreateV2AgencyRequestBodyPolicyNames
	AI_DATALAKE_AOM_POLICY           CreateV2AgencyRequestBodyPolicyNames
	AI_DATALAKE_SWR_POLICY           CreateV2AgencyRequestBodyPolicyNames
	AI_DATALAKE_LTS_POLICY           CreateV2AgencyRequestBodyPolicyNames
	AI_DATALAKE_VPCEP_POLICY         CreateV2AgencyRequestBodyPolicyNames
	AI_DATALAKE_DNS_POLICY           CreateV2AgencyRequestBodyPolicyNames
	AI_DATALAKE_OBS_POLICY           CreateV2AgencyRequestBodyPolicyNames
	AI_DATALAKE_VPC_POLICY           CreateV2AgencyRequestBodyPolicyNames
	AI_DATALAKE_DCS_POLICY           CreateV2AgencyRequestBodyPolicyNames
	AI_DATALAKE_DEW_POLICY           CreateV2AgencyRequestBodyPolicyNames
}

func GetCreateV2AgencyRequestBodyPolicyNamesEnum() CreateV2AgencyRequestBodyPolicyNamesEnum {
	return CreateV2AgencyRequestBodyPolicyNamesEnum{
		AI_DATALAKE_COMMON_POLICY: CreateV2AgencyRequestBodyPolicyNames{
			value: "AI_DATALAKE_COMMON_POLICY",
		},
		AI_DATALAKE_SELF_POLICY: CreateV2AgencyRequestBodyPolicyNames{
			value: "AI_DATALAKE_SELF_POLICY",
		},
		AI_DATALAKE_SMN_POLICY: CreateV2AgencyRequestBodyPolicyNames{
			value: "AI_DATALAKE_SMN_POLICY",
		},
		AI_DATALAKE_LAKEFORMATION_POLICY: CreateV2AgencyRequestBodyPolicyNames{
			value: "AI_DATALAKE_LAKEFORMATION_POLICY",
		},
		AI_DATALAKE_AOM_POLICY: CreateV2AgencyRequestBodyPolicyNames{
			value: "AI_DATALAKE_AOM_POLICY",
		},
		AI_DATALAKE_SWR_POLICY: CreateV2AgencyRequestBodyPolicyNames{
			value: "AI_DATALAKE_SWR_POLICY",
		},
		AI_DATALAKE_LTS_POLICY: CreateV2AgencyRequestBodyPolicyNames{
			value: "AI_DATALAKE_LTS_POLICY",
		},
		AI_DATALAKE_VPCEP_POLICY: CreateV2AgencyRequestBodyPolicyNames{
			value: "AI_DATALAKE_VPCEP_POLICY",
		},
		AI_DATALAKE_DNS_POLICY: CreateV2AgencyRequestBodyPolicyNames{
			value: "AI_DATALAKE_DNS_POLICY",
		},
		AI_DATALAKE_OBS_POLICY: CreateV2AgencyRequestBodyPolicyNames{
			value: "AI_DATALAKE_OBS_POLICY",
		},
		AI_DATALAKE_VPC_POLICY: CreateV2AgencyRequestBodyPolicyNames{
			value: "AI_DATALAKE_VPC_POLICY",
		},
		AI_DATALAKE_DCS_POLICY: CreateV2AgencyRequestBodyPolicyNames{
			value: "AI_DATALAKE_DCS_POLICY",
		},
		AI_DATALAKE_DEW_POLICY: CreateV2AgencyRequestBodyPolicyNames{
			value: "AI_DATALAKE_DEW_POLICY",
		},
	}
}

func (c CreateV2AgencyRequestBodyPolicyNames) Value() string {
	return c.value
}

func (c CreateV2AgencyRequestBodyPolicyNames) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateV2AgencyRequestBodyPolicyNames) UnmarshalJSON(b []byte) error {
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
