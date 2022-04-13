package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

// Request Object
type ShowTriggerRequest struct {
	// 消息体的类型（格式），下方类型可任选其一使用： application/json;charset=utf-8 application/json

	ContentType ShowTriggerRequestContentType `json:"Content-Type"`
	// 组织名称。小写字母开头，后面跟小写字母、数字、小数点、下划线或中划线（其中下划线最多允许连续两个，小数点、下划线、中划线不能直接相连），小写字母或数字结尾，1-64个字符。

	Namespace string `json:"namespace"`
	// 镜像仓库名称

	Repository string `json:"repository"`
	// 触发器名称

	Trigger string `json:"trigger"`
}

func (o ShowTriggerRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTriggerRequest struct{}"
	}

	return strings.Join([]string{"ShowTriggerRequest", string(data)}, " ")
}

type ShowTriggerRequestContentType struct {
	value string
}

type ShowTriggerRequestContentTypeEnum struct {
	APPLICATION_JSONCHARSETUTF_8 ShowTriggerRequestContentType
	APPLICATION_JSON             ShowTriggerRequestContentType
}

func GetShowTriggerRequestContentTypeEnum() ShowTriggerRequestContentTypeEnum {
	return ShowTriggerRequestContentTypeEnum{
		APPLICATION_JSONCHARSETUTF_8: ShowTriggerRequestContentType{
			value: "application/json;charset=utf-8",
		},
		APPLICATION_JSON: ShowTriggerRequestContentType{
			value: "application/json",
		},
	}
}

func (c ShowTriggerRequestContentType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowTriggerRequestContentType) UnmarshalJSON(b []byte) error {
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
