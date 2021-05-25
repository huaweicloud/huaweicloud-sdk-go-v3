package model

import (
	"encoding/json"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type DeleteTriggerRequest struct {
	// 消息体的类型（格式），下方类型可任选其一使用： application/json;charset=utf-8 application/json

	ContentType DeleteTriggerRequestContentType `json:"Content-Type"`
	// 组织名称

	Namespace string `json:"namespace"`
	// 镜像仓库名称

	Repository string `json:"repository"`
	// 触发器名称

	Trigger string `json:"trigger"`
}

func (o DeleteTriggerRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteTriggerRequest struct{}"
	}

	return strings.Join([]string{"DeleteTriggerRequest", string(data)}, " ")
}

type DeleteTriggerRequestContentType struct {
	value string
}

type DeleteTriggerRequestContentTypeEnum struct {
	APPLICATION_JSONCHARSETUTF_8 DeleteTriggerRequestContentType
	APPLICATION_JSON             DeleteTriggerRequestContentType
}

func GetDeleteTriggerRequestContentTypeEnum() DeleteTriggerRequestContentTypeEnum {
	return DeleteTriggerRequestContentTypeEnum{
		APPLICATION_JSONCHARSETUTF_8: DeleteTriggerRequestContentType{
			value: "application/json;charset=utf-8",
		},
		APPLICATION_JSON: DeleteTriggerRequestContentType{
			value: "application/json",
		},
	}
}

func (c DeleteTriggerRequestContentType) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *DeleteTriggerRequestContentType) UnmarshalJSON(b []byte) error {
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
