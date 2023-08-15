package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreatePostPaidInstanceByEngineRequest Request Object
type CreatePostPaidInstanceByEngineRequest struct {

	// 消息引擎。
	Engine CreatePostPaidInstanceByEngineRequestEngine `json:"engine"`

	Body *CreateInstanceReq `json:"body,omitempty"`
}

func (o CreatePostPaidInstanceByEngineRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePostPaidInstanceByEngineRequest struct{}"
	}

	return strings.Join([]string{"CreatePostPaidInstanceByEngineRequest", string(data)}, " ")
}

type CreatePostPaidInstanceByEngineRequestEngine struct {
	value string
}

type CreatePostPaidInstanceByEngineRequestEngineEnum struct {
	RABBITMQ CreatePostPaidInstanceByEngineRequestEngine
}

func GetCreatePostPaidInstanceByEngineRequestEngineEnum() CreatePostPaidInstanceByEngineRequestEngineEnum {
	return CreatePostPaidInstanceByEngineRequestEngineEnum{
		RABBITMQ: CreatePostPaidInstanceByEngineRequestEngine{
			value: "rabbitmq",
		},
	}
}

func (c CreatePostPaidInstanceByEngineRequestEngine) Value() string {
	return c.value
}

func (c CreatePostPaidInstanceByEngineRequestEngine) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreatePostPaidInstanceByEngineRequestEngine) UnmarshalJSON(b []byte) error {
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
