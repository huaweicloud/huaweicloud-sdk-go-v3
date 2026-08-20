package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ChangeDeploymentSolutionRequestBody **参数解释**: 部署方案变更请求体。 **约束限制**: 不涉及。
type ChangeDeploymentSolutionRequestBody struct {

	// **参数解释**: 变更后的目标部署形态。 **约束限制**: 必填。取值需为当前实例允许变更的目标形态。 **取值范围**: - logger：一主一备一日志节点 - triset：一主两备三节点  **默认取值**: 不涉及。
	Solution ChangeDeploymentSolutionRequestBodySolution `json:"solution"`

	// **参数解释**: 部署可用区，多个可用区以英文逗号\",\"隔开。 **约束限制**: 必填。不可包含 []()^%&\\\\'`|\";=?$<> 等特殊字符。 **取值范围**: 不涉及。 **默认取值**: 不涉及。
	AvailabilityZone string `json:"availability_zone"`

	// **参数解释**: 主可用区。 **约束限制**: 不填时默认使用当前主可用区。 **取值范围**: 不涉及。 **默认取值**: 当前实例的主可用区。
	MasterAz *string `json:"master_az,omitempty"`

	// **参数解释**: 日志可用区。 **约束限制**: 带日志节点的部署形态（如 logger）需要传该参数。 **取值范围**: 不涉及。 **默认取值**: 不涉及。
	LoggerAz *string `json:"logger_az,omitempty"`
}

func (o ChangeDeploymentSolutionRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeDeploymentSolutionRequestBody struct{}"
	}

	return strings.Join([]string{"ChangeDeploymentSolutionRequestBody", string(data)}, " ")
}

type ChangeDeploymentSolutionRequestBodySolution struct {
	value string
}

type ChangeDeploymentSolutionRequestBodySolutionEnum struct {
	LOGGER ChangeDeploymentSolutionRequestBodySolution
	TRISET ChangeDeploymentSolutionRequestBodySolution
}

func GetChangeDeploymentSolutionRequestBodySolutionEnum() ChangeDeploymentSolutionRequestBodySolutionEnum {
	return ChangeDeploymentSolutionRequestBodySolutionEnum{
		LOGGER: ChangeDeploymentSolutionRequestBodySolution{
			value: "logger",
		},
		TRISET: ChangeDeploymentSolutionRequestBodySolution{
			value: "triset",
		},
	}
}

func (c ChangeDeploymentSolutionRequestBodySolution) Value() string {
	return c.value
}

func (c ChangeDeploymentSolutionRequestBodySolution) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ChangeDeploymentSolutionRequestBodySolution) UnmarshalJSON(b []byte) error {
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
