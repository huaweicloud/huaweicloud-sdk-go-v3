package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

type ApiActionInfo struct {
	// 需要进行的操作。 - online：发布 - offline：下线

	Action ApiActionInfoAction `json:"action"`
	// API的编号，即：需要进行发布或下线的API的编号

	ApiId string `json:"api_id"`
	// 环境的编号，即：API需要发布到哪个环境

	EnvId string `json:"env_id"`
	// 对发布动作的简述。字符长度不超过255 > 中文字符必须为UTF-8或者unicode编码。

	Remark *string `json:"remark,omitempty"`
}

func (o ApiActionInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApiActionInfo struct{}"
	}

	return strings.Join([]string{"ApiActionInfo", string(data)}, " ")
}

type ApiActionInfoAction struct {
	value string
}

type ApiActionInfoActionEnum struct {
	ONLINE  ApiActionInfoAction
	OFFLINE ApiActionInfoAction
}

func GetApiActionInfoActionEnum() ApiActionInfoActionEnum {
	return ApiActionInfoActionEnum{
		ONLINE: ApiActionInfoAction{
			value: "online",
		},
		OFFLINE: ApiActionInfoAction{
			value: "offline",
		},
	}
}

func (c ApiActionInfoAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ApiActionInfoAction) UnmarshalJSON(b []byte) error {
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
