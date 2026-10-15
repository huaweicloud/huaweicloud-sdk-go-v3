package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// PublicIpConfig 租户公网IP信息
type PublicIpConfig struct {

	// **参数解释**：  指定公网IP的ID。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	Id string `json:"id"`

	// **参数解释**：  指定公网IP地址。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	PublicIp string `json:"public_ip"`

	// **参数解释**：  指定公网IP绑定任务的类型。  **约束限制**：  不涉及。  **取值范围**：  - master：主任务或单任务的公网IP - slave：备任务的公网IP  **默认取值**：  不涉及
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
