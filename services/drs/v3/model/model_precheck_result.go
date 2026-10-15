package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// PrecheckResult 预检查结果信息体
type PrecheckResult struct {

	// **参数解释**：  预检查的具体检查项名称，用于标识预检查中执行的具体检查项目。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	Item *string `json:"item,omitempty"`

	// **参数解释**：  预检查项的检查结果，表示该检查项是否通过。  **约束限制**：  不涉及。  **取值范围**：  - PASSED：通过 - ALARM：告警 - FAILED：失败  **默认取值**：  不涉及
	Result *PrecheckResultResult `json:"result,omitempty"`

	// **参数解释**：  预检查项的失败原因，当检查结果为FAILED或ALARM时返回的具体失败描述。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	FailedReason *string `json:"failed_reason,omitempty"`

	// **参数解释**：  预检查项的加密数据，包含检查项的详细检测数据，以JSON格式加密返回。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	Data *string `json:"data,omitempty"`

	// **参数解释**：  原始错误信息，预检查过程中产生的原始错误描述。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	RawErrorMsg *string `json:"raw_error_msg,omitempty"`

	// **参数解释**：  检查项分组，用于将预检查项按类别归类展示。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	Group *string `json:"group,omitempty"`

	// **参数解释**：  失败的子任务信息集合，包含预检查中失败的子任务详情，详情请参见PrecheckFailSubJobVO对象定义。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	FailedSubJobs *[]PrecheckFailSubJobVo `json:"failed_sub_jobs,omitempty"`
}

func (o PrecheckResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PrecheckResult struct{}"
	}

	return strings.Join([]string{"PrecheckResult", string(data)}, " ")
}

type PrecheckResultResult struct {
	value string
}

type PrecheckResultResultEnum struct {
	PASSED PrecheckResultResult
	ALARM  PrecheckResultResult
	FAILED PrecheckResultResult
}

func GetPrecheckResultResultEnum() PrecheckResultResultEnum {
	return PrecheckResultResultEnum{
		PASSED: PrecheckResultResult{
			value: "PASSED",
		},
		ALARM: PrecheckResultResult{
			value: "ALARM",
		},
		FAILED: PrecheckResultResult{
			value: "FAILED",
		},
	}
}

func (c PrecheckResultResult) Value() string {
	return c.value
}

func (c PrecheckResultResult) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PrecheckResultResult) UnmarshalJSON(b []byte) error {
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
