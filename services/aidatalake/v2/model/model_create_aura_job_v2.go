package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type CreateAuraJobV2 struct {

	// **参数解释**：引擎类型。 **约束限制**：不涉及。 **取值范围**：AuraJobV2。 **默认取值**：不涉及。
	Type *CreateAuraJobV2Type `json:"type,omitempty"`

	// **参数解释**：开启后，通算工作组作业能够使用智算工作组空闲CPU及内存资源，该操作可能会导致智算任务因资源不足而排队。 **约束限制**：不涉及。 **取值范围**：   - true：开启通算工作组作业使用智算工作组空闲CPU及内存资源。   - false：关闭通算工作组作业使用智算工作组空闲CPU及内存资源。 **默认取值**：false。
	ExpandCpuResource *bool `json:"expand_cpu_resource,omitempty"`

	// **参数解释**：工作组补充配置，和外部capacity对应。 **约束限制**：最多10个。
	WorkerGroupConfig []WorkerGroupConfig `json:"worker_group_config"`

	StagingConfig *AuraV2StagingConfig `json:"staging_config"`

	StoragesConfig *StoragesConfig `json:"storages_config"`
}

func (o CreateAuraJobV2) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAuraJobV2 struct{}"
	}

	return strings.Join([]string{"CreateAuraJobV2", string(data)}, " ")
}

type CreateAuraJobV2Type struct {
	value string
}

type CreateAuraJobV2TypeEnum struct {
	AURA_JOB_V2 CreateAuraJobV2Type
}

func GetCreateAuraJobV2TypeEnum() CreateAuraJobV2TypeEnum {
	return CreateAuraJobV2TypeEnum{
		AURA_JOB_V2: CreateAuraJobV2Type{
			value: "AuraJobV2",
		},
	}
}

func (c CreateAuraJobV2Type) Value() string {
	return c.value
}

func (c CreateAuraJobV2Type) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateAuraJobV2Type) UnmarshalJSON(b []byte) error {
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
