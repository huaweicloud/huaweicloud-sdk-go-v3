package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AttachV2AgencyPolicyRequestBody **参数解释**：附加策略请求。 **约束限制**：policy_name为必填字段。
type AttachV2AgencyPolicyRequestBody struct {

	// **参数解释**：权限策略名称，指定要操作的权限策略。 **约束限制**：不支持操作必选策略（AI_DATALAKE_COMMON_POLICY、AI_DATALAKE_SELF_POLICY、AI_DATALAKE_LAKEFORMATION_POLICY、AI_DATALAKE_LTS_POLICY）。 **取值范围**：   - AI_DATALAKE_SMN_POLICY：消息通知功能相关权限策略，用来将系统通知消息转发到SMN。   - AI_DATALAKE_AOM_POLICY：AOM功能相关权限策略。   - AI_DATALAKE_SWR_POLICY：镜像包注册功能相关权限策略。   - AI_DATALAKE_VPCEP_POLICY：VPCEP相关权限策略。   - AI_DATALAKE_DNS_POLICY：DNS相关权限策略。   - AI_DATALAKE_OBS_POLICY：OBS相关权限策略。   - AI_DATALAKE_VPC_POLICY：VPC相关权限策略。   - AI_DATALAKE_DCS_POLICY：DCS相关权限策略。   - AI_DATALAKE_DEW_POLICY：DEW相关权限策略。 **默认取值**：不涉及。
	PolicyName string `json:"policy_name"`
}

func (o AttachV2AgencyPolicyRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AttachV2AgencyPolicyRequestBody struct{}"
	}

	return strings.Join([]string{"AttachV2AgencyPolicyRequestBody", string(data)}, " ")
}
