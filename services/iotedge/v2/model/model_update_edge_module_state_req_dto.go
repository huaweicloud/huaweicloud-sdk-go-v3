package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// UpdateEdgeModuleStateReqDto 更新边缘模块状态请求结构体
type UpdateEdgeModuleStateReqDto struct {

	// 模块状态
	State *UpdateEdgeModuleStateReqDtoState `json:"state,omitempty"`
}

func (o UpdateEdgeModuleStateReqDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateEdgeModuleStateReqDto struct{}"
	}

	return strings.Join([]string{"UpdateEdgeModuleStateReqDto", string(data)}, " ")
}

type UpdateEdgeModuleStateReqDtoState struct {
	value string
}

type UpdateEdgeModuleStateReqDtoStateEnum struct {
	RUNNING UpdateEdgeModuleStateReqDtoState
	STOPPED UpdateEdgeModuleStateReqDtoState
}

func GetUpdateEdgeModuleStateReqDtoStateEnum() UpdateEdgeModuleStateReqDtoStateEnum {
	return UpdateEdgeModuleStateReqDtoStateEnum{
		RUNNING: UpdateEdgeModuleStateReqDtoState{
			value: "RUNNING",
		},
		STOPPED: UpdateEdgeModuleStateReqDtoState{
			value: "STOPPED",
		},
	}
}

func (c UpdateEdgeModuleStateReqDtoState) Value() string {
	return c.value
}

func (c UpdateEdgeModuleStateReqDtoState) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateEdgeModuleStateReqDtoState) UnmarshalJSON(b []byte) error {
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
