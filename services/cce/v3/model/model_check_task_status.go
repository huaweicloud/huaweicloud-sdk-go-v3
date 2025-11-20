package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type CheckTaskStatus struct {

	// **参数解释：** 插件检查状态 **取值范围：** - Init: 插件检查状态，初始化 - Running: 插件检查状态，检查中 - Failed: 插件检查状态，检查完成有风险 - Success: 插件检查状态，检查完成无风险
	Status CheckTaskStatusStatus `json:"status"`

	// **参数解释：** 插件检查结果信息 **取值范围：** 不涉及
	Message string `json:"message"`

	// **参数解释：** 插件检查风险项列表，不同插件对应的风险检查项不同。 **约束限制：** 不涉及
	RiskList *[]CheckTaskRisk `json:"riskList,omitempty"`
}

func (o CheckTaskStatus) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckTaskStatus struct{}"
	}

	return strings.Join([]string{"CheckTaskStatus", string(data)}, " ")
}

type CheckTaskStatusStatus struct {
	value string
}

type CheckTaskStatusStatusEnum struct {
	INIT    CheckTaskStatusStatus
	RUNNING CheckTaskStatusStatus
	FAILED  CheckTaskStatusStatus
	SUCCESS CheckTaskStatusStatus
}

func GetCheckTaskStatusStatusEnum() CheckTaskStatusStatusEnum {
	return CheckTaskStatusStatusEnum{
		INIT: CheckTaskStatusStatus{
			value: "Init",
		},
		RUNNING: CheckTaskStatusStatus{
			value: "Running",
		},
		FAILED: CheckTaskStatusStatus{
			value: "Failed",
		},
		SUCCESS: CheckTaskStatusStatus{
			value: "Success",
		},
	}
}

func (c CheckTaskStatusStatus) Value() string {
	return c.value
}

func (c CheckTaskStatusStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CheckTaskStatusStatus) UnmarshalJSON(b []byte) error {
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
