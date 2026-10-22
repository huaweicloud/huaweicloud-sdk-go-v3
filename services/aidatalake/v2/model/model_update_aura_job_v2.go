package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type UpdateAuraJobV2 struct {

	// **参数解释**：引擎类型。 **约束限制**：不涉及。 **取值范围**：AuraJobV2。 **默认取值**：不涉及。
	Type *UpdateAuraJobV2Type `json:"type,omitempty"`

	// **参数解释**：工作组补充配置，和外部capacity对应。 **约束限制**：最多10个。
	WorkerGroupConfig *[]WorkerGroupConfig `json:"worker_group_config,omitempty"`

	// **参数解释**：开启后，通算工作组作业能够使用智算工作组空闲CPU及内存资源，该操作可能会导致智算任务因资源不足而排队。 **约束限制**：不涉及。 **取值范围**：   - true：开启通算工作组作业使用智算工作组空闲CPU及内存资源。   - false：关闭通算工作组作业使用智算工作组空闲CPU及内存资源。 **默认取值**：false。
	ExpandCpuResource *bool `json:"expand_cpu_resource,omitempty"`
}

func (o UpdateAuraJobV2) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAuraJobV2 struct{}"
	}

	return strings.Join([]string{"UpdateAuraJobV2", string(data)}, " ")
}

type UpdateAuraJobV2Type struct {
	value string
}

type UpdateAuraJobV2TypeEnum struct {
	AURA_JOB_V2 UpdateAuraJobV2Type
}

func GetUpdateAuraJobV2TypeEnum() UpdateAuraJobV2TypeEnum {
	return UpdateAuraJobV2TypeEnum{
		AURA_JOB_V2: UpdateAuraJobV2Type{
			value: "AuraJobV2",
		},
	}
}

func (c UpdateAuraJobV2Type) Value() string {
	return c.value
}

func (c UpdateAuraJobV2Type) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateAuraJobV2Type) UnmarshalJSON(b []byte) error {
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
