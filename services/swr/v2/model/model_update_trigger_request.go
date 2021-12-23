package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type UpdateTriggerRequest struct {
	// 消息体的类型（格式），下方类型可任选其一使用： application/json;charset=utf-8 application/json

	ContentType UpdateTriggerRequestContentType `json:"Content-Type"`
	// 组织名称。小写字母开头，后面跟小写字母、数字、小数点、下划线或中划线（其中下划线最多允许连续两个，小数点、下划线、中划线不能直接相连），小写字母或数字结尾，1-64个字符。

	Namespace string `json:"namespace"`
	// 镜像仓库名称

	Repository string `json:"repository"`
	// 触发器名称

	Trigger string `json:"trigger"`

	Body *UpdateTriggerRequestBody `json:"body,omitempty"`
}

func (o UpdateTriggerRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTriggerRequest struct{}"
	}

	return strings.Join([]string{"UpdateTriggerRequest", string(data)}, " ")
}

type UpdateTriggerRequestContentType struct {
	value string
}

type UpdateTriggerRequestContentTypeEnum struct {
	APPLICATION_JSONCHARSETUTF_8 UpdateTriggerRequestContentType
	APPLICATION_JSON             UpdateTriggerRequestContentType
}

func GetUpdateTriggerRequestContentTypeEnum() UpdateTriggerRequestContentTypeEnum {
	return UpdateTriggerRequestContentTypeEnum{
		APPLICATION_JSONCHARSETUTF_8: UpdateTriggerRequestContentType{
			value: "application/json;charset=utf-8",
		},
		APPLICATION_JSON: UpdateTriggerRequestContentType{
			value: "application/json",
		},
	}
}

func (c UpdateTriggerRequestContentType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateTriggerRequestContentType) UnmarshalJSON(b []byte) error {
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
