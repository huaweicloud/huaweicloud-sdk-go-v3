package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ConstructReq 搭建容灾请求参数。
type ConstructReq struct {

	// 容灾类型。
	DisasterType ConstructReqDisasterType `json:"disaster_type"`

	// 对端实例任意数据ip。
	DrIp string `json:"dr_ip"`

	// 对端实例账户名称。
	DrUserName string `json:"dr_user_name"`

	// 对端实例账户密码。
	DrUserPassword string `json:"dr_user_password"`

	// 容灾任务名称
	DrTaskName *string `json:"dr_task_name,omitempty"`
}

func (o ConstructReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConstructReq struct{}"
	}

	return strings.Join([]string{"ConstructReq", string(data)}, " ")
}

type ConstructReqDisasterType struct {
	value string
}

type ConstructReqDisasterTypeEnum struct {
	STREAM ConstructReqDisasterType
	DORADO ConstructReqDisasterType
}

func GetConstructReqDisasterTypeEnum() ConstructReqDisasterTypeEnum {
	return ConstructReqDisasterTypeEnum{
		STREAM: ConstructReqDisasterType{
			value: "stream",
		},
		DORADO: ConstructReqDisasterType{
			value: "dorado",
		},
	}
}

func (c ConstructReqDisasterType) Value() string {
	return c.value
}

func (c ConstructReqDisasterType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ConstructReqDisasterType) UnmarshalJSON(b []byte) error {
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
