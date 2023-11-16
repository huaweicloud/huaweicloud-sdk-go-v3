package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// BlackWhiteIpRequestBody 黑白名单请求体
type BlackWhiteIpRequestBody struct {

	// 类型。white：白名单，black：黑名单
	Type BlackWhiteIpRequestBodyType `json:"type"`

	// ip列表
	IpList []string `json:"ip_list"`
}

func (o BlackWhiteIpRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BlackWhiteIpRequestBody struct{}"
	}

	return strings.Join([]string{"BlackWhiteIpRequestBody", string(data)}, " ")
}

type BlackWhiteIpRequestBodyType struct {
	value string
}

type BlackWhiteIpRequestBodyTypeEnum struct {
	WHITE BlackWhiteIpRequestBodyType
	BLACK BlackWhiteIpRequestBodyType
}

func GetBlackWhiteIpRequestBodyTypeEnum() BlackWhiteIpRequestBodyTypeEnum {
	return BlackWhiteIpRequestBodyTypeEnum{
		WHITE: BlackWhiteIpRequestBodyType{
			value: "white",
		},
		BLACK: BlackWhiteIpRequestBodyType{
			value: "black",
		},
	}
}

func (c BlackWhiteIpRequestBodyType) Value() string {
	return c.value
}

func (c BlackWhiteIpRequestBodyType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BlackWhiteIpRequestBodyType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}
