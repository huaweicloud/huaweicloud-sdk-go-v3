package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// PublicIpConfig 租户公网IP信息
type PublicIpConfig struct {

	// 指定公网IP的ID
	Id string `json:"id"`

	// 指定公网IP
	PublicIp string `json:"public_ip"`

	// 指定公网IP绑定任务的类型： - 主备任务：主是master，备是slave - 其他固定传master
	Type PublicIpConfigType `json:"type"`
}

func (o PublicIpConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PublicIpConfig struct{}"
	}

	return strings.Join([]string{"PublicIpConfig", string(data)}, " ")
}

type PublicIpConfigType struct {
	value string
}

type PublicIpConfigTypeEnum struct {
	MASTER PublicIpConfigType
	SLAVE  PublicIpConfigType
}

func GetPublicIpConfigTypeEnum() PublicIpConfigTypeEnum {
	return PublicIpConfigTypeEnum{
		MASTER: PublicIpConfigType{
			value: "master",
		},
		SLAVE: PublicIpConfigType{
			value: "slave",
		},
	}
}

func (c PublicIpConfigType) Value() string {
	return c.value
}

func (c PublicIpConfigType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PublicIpConfigType) UnmarshalJSON(b []byte) error {
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
