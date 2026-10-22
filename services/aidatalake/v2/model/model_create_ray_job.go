package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateRayJob **参数解释**：创建RayJob配置。
type CreateRayJob struct {

	// **参数解释**：引擎类型。 **约束限制**：不涉及。 **取值范围**：RayJob。 **默认取值**：不涉及。
	Type *CreateRayJobType `json:"type,omitempty"`

	HistoryServerConfig *CreateHistoryServerConfig `json:"history_server_config,omitempty"`
}

func (o CreateRayJob) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateRayJob struct{}"
	}

	return strings.Join([]string{"CreateRayJob", string(data)}, " ")
}

type CreateRayJobType struct {
	value string
}

type CreateRayJobTypeEnum struct {
	RAY_JOB CreateRayJobType
}

func GetCreateRayJobTypeEnum() CreateRayJobTypeEnum {
	return CreateRayJobTypeEnum{
		RAY_JOB: CreateRayJobType{
			value: "RayJob",
		},
	}
}

func (c CreateRayJobType) Value() string {
	return c.value
}

func (c CreateRayJobType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateRayJobType) UnmarshalJSON(b []byte) error {
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
