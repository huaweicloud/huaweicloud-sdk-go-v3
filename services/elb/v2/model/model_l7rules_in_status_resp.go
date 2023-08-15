package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// L7rulesInStatusResp 转发规则对象，用于状态树中
type L7rulesInStatusResp struct {

	// 转发规则的匹配内容。PATH：匹配请求中的路径；HOST_NAME：匹配请求中的域名
	Type L7rulesInStatusRespType `json:"type"`

	// 转发规则ID
	Id string `json:"id"`

	// 转发规则的配置状态；该字段为预留字段，暂未启用。默认为ACTIVE。
	ProvisioningStatus string `json:"provisioning_status"`
}

func (o L7rulesInStatusResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "L7rulesInStatusResp struct{}"
	}

	return strings.Join([]string{"L7rulesInStatusResp", string(data)}, " ")
}

type L7rulesInStatusRespType struct {
	value string
}

type L7rulesInStatusRespTypeEnum struct {
	PATH      L7rulesInStatusRespType
	HOST_NAME L7rulesInStatusRespType
}

func GetL7rulesInStatusRespTypeEnum() L7rulesInStatusRespTypeEnum {
	return L7rulesInStatusRespTypeEnum{
		PATH: L7rulesInStatusRespType{
			value: "PATH",
		},
		HOST_NAME: L7rulesInStatusRespType{
			value: "HOST_NAME",
		},
	}
}

func (c L7rulesInStatusRespType) Value() string {
	return c.value
}

func (c L7rulesInStatusRespType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *L7rulesInStatusRespType) UnmarshalJSON(b []byte) error {
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
