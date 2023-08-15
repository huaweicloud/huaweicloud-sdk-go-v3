package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowApiVersionRequest Request Object
type ShowApiVersionRequest struct {

	// API版本ID。
	ApiVersion ShowApiVersionRequestApiVersion `json:"api_version"`
}

func (o ShowApiVersionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowApiVersionRequest struct{}"
	}

	return strings.Join([]string{"ShowApiVersionRequest", string(data)}, " ")
}

type ShowApiVersionRequestApiVersion struct {
	value string
}

type ShowApiVersionRequestApiVersionEnum struct {
	V1 ShowApiVersionRequestApiVersion
	V2 ShowApiVersionRequestApiVersion
}

func GetShowApiVersionRequestApiVersionEnum() ShowApiVersionRequestApiVersionEnum {
	return ShowApiVersionRequestApiVersionEnum{
		V1: ShowApiVersionRequestApiVersion{
			value: "v1",
		},
		V2: ShowApiVersionRequestApiVersion{
			value: "v2",
		},
	}
}

func (c ShowApiVersionRequestApiVersion) Value() string {
	return c.value
}

func (c ShowApiVersionRequestApiVersion) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowApiVersionRequestApiVersion) UnmarshalJSON(b []byte) error {
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
