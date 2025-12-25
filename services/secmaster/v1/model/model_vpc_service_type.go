package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// VpcServiceType **参数解释**: Vpc服务状态 - MANAGE 管理通道 - DATA 数据通道  **约束限制** 不涉及  **取值范围**: - MANAGE - DATA  **默认值** 不涉及
type VpcServiceType struct {
	value string
}

type VpcServiceTypeEnum struct {
	MANAGE VpcServiceType
	DATA   VpcServiceType
}

func GetVpcServiceTypeEnum() VpcServiceTypeEnum {
	return VpcServiceTypeEnum{
		MANAGE: VpcServiceType{
			value: "MANAGE",
		},
		DATA: VpcServiceType{
			value: "DATA",
		},
	}
}

func (c VpcServiceType) Value() string {
	return c.value
}

func (c VpcServiceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *VpcServiceType) UnmarshalJSON(b []byte) error {
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
