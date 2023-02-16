package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type CreateInstanceByEngineRequest struct {

	// 消息引擎。
	Engine CreateInstanceByEngineRequestEngine `json:"engine"`

	Body *CreateInstanceByEngineReq `json:"body,omitempty"`
}

func (o CreateInstanceByEngineRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateInstanceByEngineRequest struct{}"
	}

	return strings.Join([]string{"CreateInstanceByEngineRequest", string(data)}, " ")
}

type CreateInstanceByEngineRequestEngine struct {
	value string
}

type CreateInstanceByEngineRequestEngineEnum struct {
	RELIABILITY CreateInstanceByEngineRequestEngine
}

func GetCreateInstanceByEngineRequestEngineEnum() CreateInstanceByEngineRequestEngineEnum {
	return CreateInstanceByEngineRequestEngineEnum{
		RELIABILITY: CreateInstanceByEngineRequestEngine{
			value: "reliability",
		},
	}
}

func (c CreateInstanceByEngineRequestEngine) Value() string {
	return c.value
}

func (c CreateInstanceByEngineRequestEngine) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateInstanceByEngineRequestEngine) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
