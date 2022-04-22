package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type ShowApiVersionRequest struct {

	// 消息体的类型（格式），下方类型可任选其一使用： application/json;charset=utf-8 application/json
	ContentType ShowApiVersionRequestContentType `json:"Content-Type"`

	// API版本号。
	ApiVersion string `json:"api_version"`
}

func (o ShowApiVersionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowApiVersionRequest struct{}"
	}

	return strings.Join([]string{"ShowApiVersionRequest", string(data)}, " ")
}

type ShowApiVersionRequestContentType struct {
	value string
}

type ShowApiVersionRequestContentTypeEnum struct {
	APPLICATION_JSONCHARSETUTF_8 ShowApiVersionRequestContentType
	APPLICATION_JSON             ShowApiVersionRequestContentType
}

func GetShowApiVersionRequestContentTypeEnum() ShowApiVersionRequestContentTypeEnum {
	return ShowApiVersionRequestContentTypeEnum{
		APPLICATION_JSONCHARSETUTF_8: ShowApiVersionRequestContentType{
			value: "application/json;charset=utf-8",
		},
		APPLICATION_JSON: ShowApiVersionRequestContentType{
			value: "application/json",
		},
	}
}

func (c ShowApiVersionRequestContentType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowApiVersionRequestContentType) UnmarshalJSON(b []byte) error {
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
