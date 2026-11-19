package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ListAiInstancesRequestBody struct {

	//
	AgentName *string `json:"agent_name,omitempty"`

	// **参数解释**： 最近检测时间 **取值范围**： 不涉及
	DetectTime *int64 `json:"detect_time,omitempty"`

	// **参数解释**： 风险总数 **取值范围**： 最小值0，最大值2147483647
	SecurityRiskNum *int32 `json:"security_risk_num,omitempty"`

	// **参数解释**： 免费体检的服务器总数 **取值范围**： 最小值0，最大值2147483647
	FreeHealthCheckTotalHostNum *int32 `json:"free_health_check_total_host_num,omitempty"`

	// **参数解释**： 智能体ID **约束限制**： 不涉及 **取值范围**： 字符长度为1-128位 **默认取值**： 不涉及
	AgentId *string `json:"agent_id,omitempty"`

	// **参数解释**： 智能体ID列表，支持批量查询 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	AgentIdList *[]string `json:"agent_id_list,omitempty"`

	// **参数解释**： 智能体运行时名字 **约束限制**： 不涉及 **取值范围**： 字符长度为1-128位 **默认取值**： 不涉及
	RuntimeName *string `json:"runtime_name,omitempty"`

	// **参数解释**： 智能体运行时ID **约束限制**： 不涉及 **取值范围**： 字符长度为1-128位 **默认取值**： 不涉及
	RuntimeId *string `json:"runtime_id,omitempty"`

	// **参数解释**： 运行镜像地址 **约束限制**： 不涉及 **取值范围**： 字符长度为1-128位 **默认取值**： 不涉及
	ImageUrl *string `json:"image_url,omitempty"`

	// **参数解释**： SDK检测状态 **约束限制**： 不涉及 **取值范围**： - open：开启 - close：关闭  **默认取值**： 不涉及
	SdkStatus *string `json:"sdk_status,omitempty"`

	// **参数解释**： 是否存在行为轨迹检测 **约束限制**： 不涉及 **取值范围**： - true：存在 - false：不存在  **默认取值**： 不涉及
	ExistRiskTrace *bool `json:"exist_risk_trace,omitempty"`

	// **参数解释**： 是否存在运行时检测 **约束限制**： 不涉及 **取值范围**： - true：存在 - false：不存在  **默认取值**： 不涉及
	ExistRiskRuntime *bool `json:"exist_risk_runtime,omitempty"`

	// **参数解释**： 智能体的来源 **约束限制**： 不涉及 **取值范围**： - AgentArts：AgentArts系统。 - OfficeClaw：OfficeClaw系统。 - OpenClaw：OpenClaw系统。  **默认取值**： 不涉及
	Source *string `json:"source,omitempty"`

	// **参数解释**： 智能体类型 **约束限制**： 不涉及 **取值范围**： - agent：单智能体。 - chat：工作流。 - controller：多智能体。 - teamAgent：团队智能体。  **默认取值**： 不涉及
	AgentType *ListAiInstancesRequestBodyAgentType `json:"agent_type,omitempty"`
}

func (o ListAiInstancesRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAiInstancesRequestBody struct{}"
	}

	return strings.Join([]string{"ListAiInstancesRequestBody", string(data)}, " ")
}

type ListAiInstancesRequestBodyAgentType struct {
	value string
}

type ListAiInstancesRequestBodyAgentTypeEnum struct {
	AGENT      ListAiInstancesRequestBodyAgentType
	CHAT       ListAiInstancesRequestBodyAgentType
	CONTROLLER ListAiInstancesRequestBodyAgentType
	TEAM_AGENT ListAiInstancesRequestBodyAgentType
}

func GetListAiInstancesRequestBodyAgentTypeEnum() ListAiInstancesRequestBodyAgentTypeEnum {
	return ListAiInstancesRequestBodyAgentTypeEnum{
		AGENT: ListAiInstancesRequestBodyAgentType{
			value: "agent",
		},
		CHAT: ListAiInstancesRequestBodyAgentType{
			value: "chat",
		},
		CONTROLLER: ListAiInstancesRequestBodyAgentType{
			value: "controller",
		},
		TEAM_AGENT: ListAiInstancesRequestBodyAgentType{
			value: "teamAgent",
		},
	}
}

func (c ListAiInstancesRequestBodyAgentType) Value() string {
	return c.value
}

func (c ListAiInstancesRequestBodyAgentType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListAiInstancesRequestBodyAgentType) UnmarshalJSON(b []byte) error {
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
