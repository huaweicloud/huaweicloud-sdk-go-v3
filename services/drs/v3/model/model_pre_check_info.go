package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// PreCheckInfo 预检查信息体
type PreCheckInfo struct {

	// **参数解释**：  迁移任务ID，用于唯一标识一个数据复制任务，在查询、修改、删除等操作中作为任务的识别凭证。  **约束限制**：  不涉及。  **取值范围**：  UUID格式字符串。  **默认取值**：  不涉及
	JobId string `json:"job_id"`

	// **参数解释**：  预检查模式，指定执行预检查的场景类型。请按照任务对应的状态传值，否则可能导致预检查失败，影响任务启动。  **约束限制**：  不涉及。  **取值范围**：  - forStartJob：任务配置中首次启动时预检查 - forRetryJob：任务增量中编辑对象后预检查，或任务失败后重试时预检查  **默认取值**：  不涉及
	PrecheckMode PreCheckInfoPrecheckMode `json:"precheck_mode"`
}

func (o PreCheckInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PreCheckInfo struct{}"
	}

	return strings.Join([]string{"PreCheckInfo", string(data)}, " ")
}

type PreCheckInfoPrecheckMode struct {
	value string
}

type PreCheckInfoPrecheckModeEnum struct {
	FOR_START_JOB PreCheckInfoPrecheckMode
	FOR_RETRY_JOB PreCheckInfoPrecheckMode
}

func GetPreCheckInfoPrecheckModeEnum() PreCheckInfoPrecheckModeEnum {
	return PreCheckInfoPrecheckModeEnum{
		FOR_START_JOB: PreCheckInfoPrecheckMode{
			value: "forStartJob",
		},
		FOR_RETRY_JOB: PreCheckInfoPrecheckMode{
			value: "forRetryJob",
		},
	}
}

func (c PreCheckInfoPrecheckMode) Value() string {
	return c.value
}

func (c PreCheckInfoPrecheckMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PreCheckInfoPrecheckMode) UnmarshalJSON(b []byte) error {
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
