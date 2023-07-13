package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateRequestRequest Request Object
type CreateRequestRequest struct {

	// the number of samples
	Topn *int32 `json:"topn,omitempty"`

	// the scenario of code content
	Scenario *string `json:"scenario,omitempty"`

	// if `resubmit` is true, the de-duplication will be ignored
	Resubmit *bool `json:"resubmit,omitempty"`

	// choose the model
	ModelId *string `json:"model_id,omitempty"`

	// An enumeration. - function - rawtext
	RequestType *CreateRequestRequestRequestType `json:"request_type,omitempty"`

	Body *PropertiesSchema `json:"body,omitempty"`
}

func (o CreateRequestRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateRequestRequest struct{}"
	}

	return strings.Join([]string{"CreateRequestRequest", string(data)}, " ")
}

type CreateRequestRequestRequestType struct {
	value string
}

type CreateRequestRequestRequestTypeEnum struct {
	FUNCTION CreateRequestRequestRequestType
	RAWTEXT  CreateRequestRequestRequestType
}

func GetCreateRequestRequestRequestTypeEnum() CreateRequestRequestRequestTypeEnum {
	return CreateRequestRequestRequestTypeEnum{
		FUNCTION: CreateRequestRequestRequestType{
			value: "function",
		},
		RAWTEXT: CreateRequestRequestRequestType{
			value: "rawtext",
		},
	}
}

func (c CreateRequestRequestRequestType) Value() string {
	return c.value
}

func (c CreateRequestRequestRequestType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateRequestRequestRequestType) UnmarshalJSON(b []byte) error {
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
