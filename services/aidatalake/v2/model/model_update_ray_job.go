package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// UpdateRayJob **参数解释**：更新RayJob配置。
type UpdateRayJob struct {

	// **参数解释**：引擎类型。 **约束限制**：不涉及。 **取值范围**：RayJob。 **默认取值**：不涉及。
	Type *UpdateRayJobType `json:"type,omitempty"`

	HistoryServerConfig *HistoryServerConfig `json:"history_server_config,omitempty"`
}

func (o UpdateRayJob) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateRayJob struct{}"
	}

	return strings.Join([]string{"UpdateRayJob", string(data)}, " ")
}

type UpdateRayJobType struct {
	value string
}

type UpdateRayJobTypeEnum struct {
	RAY_JOB UpdateRayJobType
}

func GetUpdateRayJobTypeEnum() UpdateRayJobTypeEnum {
	return UpdateRayJobTypeEnum{
		RAY_JOB: UpdateRayJobType{
			value: "RayJob",
		},
	}
}

func (c UpdateRayJobType) Value() string {
	return c.value
}

func (c UpdateRayJobType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateRayJobType) UnmarshalJSON(b []byte) error {
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
