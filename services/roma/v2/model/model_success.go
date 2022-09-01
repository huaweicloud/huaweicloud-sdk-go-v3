package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type Success struct {

	// API请求路径
	Path *string `json:"path,omitempty" xml:"path"`

	// API请求方法
	Method *string `json:"method,omitempty" xml:"method"`

	// 导入行为： - update：表示更新API - create：表示新建API
	Action *SuccessAction `json:"action,omitempty" xml:"action"`

	// 导入成功的API编号
	Id *string `json:"id,omitempty" xml:"id"`
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

func (c SuccessAction) Value() string {
	return c.value
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
