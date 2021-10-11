package model

import (
	"encoding/json"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type SwaggerSuccessResp struct {
	// API请求路径

	Path *string `json:"path,omitempty"`
	// API请求方法

	Method *string `json:"method,omitempty"`
	// 导入行为： - update：表示更新API - create：表示新建API

	Action *SwaggerSuccessRespAction `json:"action,omitempty"`
	// 导入成功的API编号

	Id *string `json:"id,omitempty"`
}

func (o SwaggerSuccessResp) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "SwaggerSuccessResp struct{}"
	}

	return strings.Join([]string{"SwaggerSuccessResp", string(data)}, " ")
}

type SwaggerSuccessRespAction struct {
	value string
}

type SwaggerSuccessRespActionEnum struct {
	UPDATE SwaggerSuccessRespAction
	CREATE SwaggerSuccessRespAction
}

func GetSwaggerSuccessRespActionEnum() SwaggerSuccessRespActionEnum {
	return SwaggerSuccessRespActionEnum{
		UPDATE: SwaggerSuccessRespAction{
			value: "update",
		},
		CREATE: SwaggerSuccessRespAction{
			value: "create",
		},
	}
}

func (c SwaggerSuccessRespAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *SwaggerSuccessRespAction) UnmarshalJSON(b []byte) error {
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
