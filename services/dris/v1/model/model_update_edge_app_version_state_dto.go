package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type UpdateEdgeAppVersionStateDto struct {

	// **参数说明**：应用版本状态。  **取值范围**：  - PUBLISHED：发布  - OFF_SHELF：下线
	State *UpdateEdgeAppVersionStateDtoState `json:"state,omitempty"`
}

func (o UpdateEdgeAppVersionStateDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateEdgeAppVersionStateDto struct{}"
	}

	return strings.Join([]string{"UpdateEdgeAppVersionStateDto", string(data)}, " ")
}

type UpdateEdgeAppVersionStateDtoState struct {
	value string
}

type UpdateEdgeAppVersionStateDtoStateEnum struct {
	PUBLISHED UpdateEdgeAppVersionStateDtoState
	OFF_SHELF UpdateEdgeAppVersionStateDtoState
}

func GetUpdateEdgeAppVersionStateDtoStateEnum() UpdateEdgeAppVersionStateDtoStateEnum {
	return UpdateEdgeAppVersionStateDtoStateEnum{
		PUBLISHED: UpdateEdgeAppVersionStateDtoState{
			value: "PUBLISHED",
		},
		OFF_SHELF: UpdateEdgeAppVersionStateDtoState{
			value: "OFF_SHELF",
		},
	}
}

func (c UpdateEdgeAppVersionStateDtoState) Value() string {
	return c.value
}

func (c UpdateEdgeAppVersionStateDtoState) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateEdgeAppVersionStateDtoState) UnmarshalJSON(b []byte) error {
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
