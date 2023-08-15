package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowVersionRequest Request Object
type ShowVersionRequest struct {

	// 查询的目标版本号。 取值为：v1、v2、v3。
	Version ShowVersionRequestVersion `json:"version"`
}

func (o ShowVersionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowVersionRequest struct{}"
	}

	return strings.Join([]string{"ShowVersionRequest", string(data)}, " ")
}

type ShowVersionRequestVersion struct {
	value string
}

type ShowVersionRequestVersionEnum struct {
	V1 ShowVersionRequestVersion
	V2 ShowVersionRequestVersion
	V3 ShowVersionRequestVersion
}

func GetShowVersionRequestVersionEnum() ShowVersionRequestVersionEnum {
	return ShowVersionRequestVersionEnum{
		V1: ShowVersionRequestVersion{
			value: "v1",
		},
		V2: ShowVersionRequestVersion{
			value: "v2",
		},
		V3: ShowVersionRequestVersion{
			value: "v3",
		},
	}
}

func (c ShowVersionRequestVersion) Value() string {
	return c.value
}

func (c ShowVersionRequestVersion) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowVersionRequestVersion) UnmarshalJSON(b []byte) error {
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
