package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

// 扩展更新镜像接口请求体
type UpdateImageRequestBody struct {
	// 操作类型，目前取值为add，replace和remove。

	Op UpdateImageRequestBodyOp `json:"op"`
	// 需要更新的属性名称，需要在属性名称前加“/”。

	Path string `json:"path"`
	// 需要更新属性的值。

	Value string `json:"value"`
}

func (o UpdateImageRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateImageRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateImageRequestBody", string(data)}, " ")
}

type UpdateImageRequestBodyOp struct {
	value string
}

type UpdateImageRequestBodyOpEnum struct {
	ADD     UpdateImageRequestBodyOp
	REPLACE UpdateImageRequestBodyOp
	REMOVE  UpdateImageRequestBodyOp
}

func GetUpdateImageRequestBodyOpEnum() UpdateImageRequestBodyOpEnum {
	return UpdateImageRequestBodyOpEnum{
		ADD: UpdateImageRequestBodyOp{
			value: "add",
		},
		REPLACE: UpdateImageRequestBodyOp{
			value: "replace",
		},
		REMOVE: UpdateImageRequestBodyOp{
			value: "remove",
		},
	}
}

func (c UpdateImageRequestBodyOp) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateImageRequestBodyOp) UnmarshalJSON(b []byte) error {
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
