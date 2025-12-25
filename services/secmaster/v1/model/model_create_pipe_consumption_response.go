package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreatePipeConsumptionResponse Response Object
type CreatePipeConsumptionResponse struct {

	// 访问点地址
	AccessPoint *string `json:"access_point,omitempty"`

	// 数据空间ID
	DataspaceId *string `json:"dataspace_id,omitempty"`

	// 管道ID
	PipeId *string `json:"pipe_id,omitempty"`

	// 管道名称
	PipeName *string `json:"pipe_name,omitempty"`

	// 管道状态
	Status *CreatePipeConsumptionResponseStatus `json:"status,omitempty"`

	// 订阅ID
	Subscription *string `json:"subscription,omitempty"`

	// 管道类型
	Type           *CreatePipeConsumptionResponseType `json:"type,omitempty"`
	HttpStatusCode int                                `json:"-"`
}

func (o CreatePipeConsumptionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePipeConsumptionResponse struct{}"
	}

	return strings.Join([]string{"CreatePipeConsumptionResponse", string(data)}, " ")
}

type CreatePipeConsumptionResponseStatus struct {
	value string
}

type CreatePipeConsumptionResponseStatusEnum struct {
	ENABLE  CreatePipeConsumptionResponseStatus
	DISABLE CreatePipeConsumptionResponseStatus
}

func GetCreatePipeConsumptionResponseStatusEnum() CreatePipeConsumptionResponseStatusEnum {
	return CreatePipeConsumptionResponseStatusEnum{
		ENABLE: CreatePipeConsumptionResponseStatus{
			value: "enable",
		},
		DISABLE: CreatePipeConsumptionResponseStatus{
			value: "disable",
		},
	}
}

func (c CreatePipeConsumptionResponseStatus) Value() string {
	return c.value
}

func (c CreatePipeConsumptionResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreatePipeConsumptionResponseStatus) UnmarshalJSON(b []byte) error {
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

type CreatePipeConsumptionResponseType struct {
	value string
}

type CreatePipeConsumptionResponseTypeEnum struct {
	INTRANET CreatePipeConsumptionResponseType
	INTERNET CreatePipeConsumptionResponseType
}

func GetCreatePipeConsumptionResponseTypeEnum() CreatePipeConsumptionResponseTypeEnum {
	return CreatePipeConsumptionResponseTypeEnum{
		INTRANET: CreatePipeConsumptionResponseType{
			value: "intranet",
		},
		INTERNET: CreatePipeConsumptionResponseType{
			value: "internet",
		},
	}
}

func (c CreatePipeConsumptionResponseType) Value() string {
	return c.value
}

func (c CreatePipeConsumptionResponseType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreatePipeConsumptionResponseType) UnmarshalJSON(b []byte) error {
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
