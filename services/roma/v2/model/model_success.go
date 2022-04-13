package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

type Success struct {
	// API请求路径

	Path *string `json:"path,omitempty"`
	// API请求方法

	Method *string `json:"method,omitempty"`
	// 导入行为： - update：表示更新API - create：表示新建API

	Action *SuccessAction `json:"action,omitempty"`
	// 导入成功的API编号

	Id *string `json:"id,omitempty"`
}

func (o Success) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Success struct{}"
	}

	return strings.Join([]string{"Success", string(data)}, " ")
}

type SuccessAction struct {
	value string
}

type SuccessActionEnum struct {
	UPDATE SuccessAction
	CREATE SuccessAction
}

func GetSuccessActionEnum() SuccessActionEnum {
	return SuccessActionEnum{
		UPDATE: SuccessAction{
			value: "update",
		},
		CREATE: SuccessAction{
			value: "create",
		},
	}
}

func (c SuccessAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SuccessAction) UnmarshalJSON(b []byte) error {
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
