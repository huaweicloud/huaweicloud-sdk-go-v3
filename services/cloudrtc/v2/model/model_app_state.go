package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// AppState app状态信息
type AppState struct {

	// 状态 - ACTIVATION：开启 - DEACTIVATION：停用 - ARREARS：欠费 - DELETED：已删除
	State *AppStateState `json:"state,omitempty"`

	// app鉴权的更新时间，形如“2006-01-02T15:04:05.075Z”，时区为：UTC
	UpdateTime *string `json:"update_time,omitempty"`
}

func (o AppState) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AppState struct{}"
	}

	return strings.Join([]string{"AppState", string(data)}, " ")
}

type AppStateState struct {
	value string
}

type AppStateStateEnum struct {
	ACTIVATION   AppStateState
	DEACTIVATION AppStateState
	ARREARS      AppStateState
	DELETED      AppStateState
}

func GetAppStateStateEnum() AppStateStateEnum {
	return AppStateStateEnum{
		ACTIVATION: AppStateState{
			value: "ACTIVATION",
		},
		DEACTIVATION: AppStateState{
			value: "DEACTIVATION",
		},
		ARREARS: AppStateState{
			value: "ARREARS",
		},
		DELETED: AppStateState{
			value: "DELETED",
		},
	}
}

func (c AppStateState) Value() string {
	return c.value
}

func (c AppStateState) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AppStateState) UnmarshalJSON(b []byte) error {
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
