package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateSqlValidationResponse Response Object
type CreateSqlValidationResponse struct {

	// 源表
	Sources *[]TableItem `json:"sources,omitempty"`

	// 模式
	Modes          *[]CreateSqlValidationResponseModes `json:"modes,omitempty"`
	HttpStatusCode int                                 `json:"-"`
}

func (o CreateSqlValidationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSqlValidationResponse struct{}"
	}

	return strings.Join([]string{"CreateSqlValidationResponse", string(data)}, " ")
}

type CreateSqlValidationResponseModes struct {
	value string
}

type CreateSqlValidationResponseModesEnum struct {
	STREAMING CreateSqlValidationResponseModes
	BATCH     CreateSqlValidationResponseModes
}

func GetCreateSqlValidationResponseModesEnum() CreateSqlValidationResponseModesEnum {
	return CreateSqlValidationResponseModesEnum{
		STREAMING: CreateSqlValidationResponseModes{
			value: "STREAMING",
		},
		BATCH: CreateSqlValidationResponseModes{
			value: "BATCH",
		},
	}
}

func (c CreateSqlValidationResponseModes) Value() string {
	return c.value
}

func (c CreateSqlValidationResponseModes) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateSqlValidationResponseModes) UnmarshalJSON(b []byte) error {
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
