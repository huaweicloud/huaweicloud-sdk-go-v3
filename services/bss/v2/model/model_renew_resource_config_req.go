package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type RenewResourceConfigReq struct {

	// |参数名称：续订资源设置属性类型| |参数的约束及描述：该参数必填，设置类型，支持枚举| |DEDUCTION_DATE：设置自动续费扣款日，EXPIRE_DATE：设置续费后资源统一到期日|
	ConfigName RenewResourceConfigReqConfigName `json:"config_name"`

	// |参数名称：续订资源设置属性值| |参数约束及描述：该参数必填，当config_name值为DEDUCTION_DATE时，支持设置资源到期前2天至前7天自动扣款，config_value范围：2，3，4，5，6，7。 当config_name值为EXPIRE_DATE时，支持设置统一到期日为每个月的同一天（1~28号及每个月最后一天），config_value范围：1，2，3，4，5，6，7，8，9，10，11，12，13，14，15，16，17，18，19，20，21，22，23，24，25，26，27，28，-1|
	ConfigValue string `json:"config_value"`
}

func (o RenewResourceConfigReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RenewResourceConfigReq struct{}"
	}

	return strings.Join([]string{"RenewResourceConfigReq", string(data)}, " ")
}

type RenewResourceConfigReqConfigName struct {
	value string
}

type RenewResourceConfigReqConfigNameEnum struct {
	DEDUCTION_DATE RenewResourceConfigReqConfigName
	EXPIRE_DATE    RenewResourceConfigReqConfigName
}

func GetRenewResourceConfigReqConfigNameEnum() RenewResourceConfigReqConfigNameEnum {
	return RenewResourceConfigReqConfigNameEnum{
		DEDUCTION_DATE: RenewResourceConfigReqConfigName{
			value: "DEDUCTION_DATE",
		},
		EXPIRE_DATE: RenewResourceConfigReqConfigName{
			value: "EXPIRE_DATE",
		},
	}
}

func (c RenewResourceConfigReqConfigName) Value() string {
	return c.value
}

func (c RenewResourceConfigReqConfigName) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *RenewResourceConfigReqConfigName) UnmarshalJSON(b []byte) error {
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
