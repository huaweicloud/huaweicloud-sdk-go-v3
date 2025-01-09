package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type UserAccessStage struct {

	// 接入阶段 | LOGIN - 登录 PRECONNECT - 预连接 CONNECT - 正式连接
	Stage *UserAccessStageStage `json:"stage,omitempty"`

	// 花费时长，单位：ms
	Duration *int32 `json:"duration,omitempty"`

	// 接入阶段是否成功
	IsSuccess *bool `json:"is_success,omitempty"`

	// 开始时间戳
	StartTime *int64 `json:"start_time,omitempty"`

	// 结束时间戳
	EndTime *int64 `json:"end_time,omitempty"`

	// 错误码
	ErrorCode *string `json:"error_code,omitempty"`

	// 错误描述
	ErrorMsg *string `json:"error_msg,omitempty"`
}

func (o UserAccessStage) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UserAccessStage struct{}"
	}

	return strings.Join([]string{"UserAccessStage", string(data)}, " ")
}

type UserAccessStageStage struct {
	value string
}

type UserAccessStageStageEnum struct {
	LOGIN      UserAccessStageStage
	PRECONNECT UserAccessStageStage
	CONNECT    UserAccessStageStage
}

func GetUserAccessStageStageEnum() UserAccessStageStageEnum {
	return UserAccessStageStageEnum{
		LOGIN: UserAccessStageStage{
			value: "LOGIN",
		},
		PRECONNECT: UserAccessStageStage{
			value: "PRECONNECT",
		},
		CONNECT: UserAccessStageStage{
			value: "CONNECT",
		},
	}
}

func (c UserAccessStageStage) Value() string {
	return c.value
}

func (c UserAccessStageStage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UserAccessStageStage) UnmarshalJSON(b []byte) error {
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
