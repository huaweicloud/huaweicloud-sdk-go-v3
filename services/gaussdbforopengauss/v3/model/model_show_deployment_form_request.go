package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowDeploymentFormRequest Request Object
type ShowDeploymentFormRequest struct {

	// **参数解释**: 语言。 **约束限制**: 不涉及。 **取值范围**:   - zh-cn   - en-us  **默认取值**: en-us
	XLanguage *string `json:"X-Language,omitempty"`

	// 解决方案模板名称。
	Solution *ShowDeploymentFormRequestSolution `json:"solution,omitempty"`

	// 实例ID。
	InstanceId *string `json:"instance_id,omitempty"`

	// **参数解释**: 事务一致性类型。 **约束限制**: 不涉及。 **取值范围**: - strong - eventual **默认取值**: 不涉及。
	Consistency *string `json:"consistency,omitempty"`

	// **参数解释**: 副本一致性协议类型。 **约束限制**: 不涉及。 **取值范围**: - quorum - paxos **默认取值**: 不涉及。
	ConsistencyProtocol *string `json:"consistency_protocol,omitempty"`

	// **参数解释**: 引擎版本号。 **约束限制**: 不涉及。 **取值范围**: 不涉及。 **默认取值**: 不涉及。
	EngineVersion *string `json:"engine_version,omitempty"`
}

func (o ShowDeploymentFormRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDeploymentFormRequest struct{}"
	}

	return strings.Join([]string{"ShowDeploymentFormRequest", string(data)}, " ")
}

type ShowDeploymentFormRequestSolution struct {
	value string
}

type ShowDeploymentFormRequestSolutionEnum struct {
	TRISET ShowDeploymentFormRequestSolution
	SINGLE ShowDeploymentFormRequestSolution
}

func GetShowDeploymentFormRequestSolutionEnum() ShowDeploymentFormRequestSolutionEnum {
	return ShowDeploymentFormRequestSolutionEnum{
		TRISET: ShowDeploymentFormRequestSolution{
			value: "triset",
		},
		SINGLE: ShowDeploymentFormRequestSolution{
			value: "single",
		},
	}
}

func (c ShowDeploymentFormRequestSolution) Value() string {
	return c.value
}

func (c ShowDeploymentFormRequestSolution) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowDeploymentFormRequestSolution) UnmarshalJSON(b []byte) error {
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
