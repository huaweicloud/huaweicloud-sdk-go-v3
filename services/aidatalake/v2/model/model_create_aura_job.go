package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type CreateAuraJob struct {

	// **参数解释**：引擎类型。 **约束限制**：不涉及。 **取值范围**：AuraJob。 **默认取值**：不涉及。
	Type *CreateAuraJobType `json:"type,omitempty"`

	StoragesConfig *StoragesConfig `json:"storages_config,omitempty"`
}

func (o CreateAuraJob) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAuraJob struct{}"
	}

	return strings.Join([]string{"CreateAuraJob", string(data)}, " ")
}

type CreateAuraJobType struct {
	value string
}

type CreateAuraJobTypeEnum struct {
	AURA_JOB CreateAuraJobType
}

func GetCreateAuraJobTypeEnum() CreateAuraJobTypeEnum {
	return CreateAuraJobTypeEnum{
		AURA_JOB: CreateAuraJobType{
			value: "AuraJob",
		},
	}
}

func (c CreateAuraJobType) Value() string {
	return c.value
}

func (c CreateAuraJobType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateAuraJobType) UnmarshalJSON(b []byte) error {
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
