package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CheckJobNameReq 校验任务名称请求体。
type CheckJobNameReq struct {

	// - 迁移、同步、灾备、订阅任务名称，名称在4位到50位之间，必须以字母开头，可以包含字母、数字、中划线或下划线，不能包含其他特殊字符，任务名称不能重复。 - 备份迁移任务名称，名称在4位到80位之间，必须以字母开头，可以包含字母、数字、中划线或下划线，不能包含其他特殊字符，任务名称不能重复。
	Name string `json:"name"`

	// 任务类型。 - trans - subscription - offline
	Type CheckJobNameReqType `json:"type"`
}

func (o CheckJobNameReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckJobNameReq struct{}"
	}

	return strings.Join([]string{"CheckJobNameReq", string(data)}, " ")
}

type CheckJobNameReqType struct {
	value string
}

type CheckJobNameReqTypeEnum struct {
	TRANS        CheckJobNameReqType
	SUBSCRIPTION CheckJobNameReqType
	OFFLINE      CheckJobNameReqType
}

func GetCheckJobNameReqTypeEnum() CheckJobNameReqTypeEnum {
	return CheckJobNameReqTypeEnum{
		TRANS: CheckJobNameReqType{
			value: "trans",
		},
		SUBSCRIPTION: CheckJobNameReqType{
			value: "subscription",
		},
		OFFLINE: CheckJobNameReqType{
			value: "offline",
		},
	}
}

func (c CheckJobNameReqType) Value() string {
	return c.value
}

func (c CheckJobNameReqType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CheckJobNameReqType) UnmarshalJSON(b []byte) error {
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
