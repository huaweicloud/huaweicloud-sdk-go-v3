package model

import (
	"encoding/json"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type CreateTriggerRequest struct {
	// 消息体的类型（格式），下方类型可任选其一使用： application/json;charset=utf-8 application/json

	ContentType CreateTriggerRequestContentType `json:"Content-Type"`
	// 组织名称

	Namespace string `json:"namespace"`
	// 镜像仓库名称

	Repository string `json:"repository"`

	Body *CreateTriggerRequestBody `json:"body,omitempty"`
}

func (o CreateTriggerRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateTriggerRequest struct{}"
	}

	return strings.Join([]string{"CreateTriggerRequest", string(data)}, " ")
}

type CreateTriggerRequestContentType struct {
	value string
}

type CreateTriggerRequestContentTypeEnum struct {
	APPLICATION_JSONCHARSETUTF_8 CreateTriggerRequestContentType
	APPLICATION_JSON             CreateTriggerRequestContentType
}

func GetCreateTriggerRequestContentTypeEnum() CreateTriggerRequestContentTypeEnum {
	return CreateTriggerRequestContentTypeEnum{
		APPLICATION_JSONCHARSETUTF_8: CreateTriggerRequestContentType{
			value: "application/json;charset=utf-8",
		},
		APPLICATION_JSON: CreateTriggerRequestContentType{
			value: "application/json",
		},
	}
}

func (c CreateTriggerRequestContentType) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *CreateTriggerRequestContentType) UnmarshalJSON(b []byte) error {
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
