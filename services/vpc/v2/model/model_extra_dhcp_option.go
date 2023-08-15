package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ExtraDhcpOption 子网配置的NTP地址或租约时间对象
type ExtraDhcpOption struct {

	// 功能说明：子网配置的NTP地址名称或子网配置的租约到期名称。 约束：目前只支持字段“ntp”或“addresstime”
	OptName ExtraDhcpOptionOptName `json:"opt_name"`

	// 功能说明：子网配置的NTP地址或子网配置的租约到期时间。 约束：opt_name配置为“ntp”，则表示是子网ntp地址，目前只支持IPv4地址，每个IP地址以逗号隔开，IP地址个数不能超过4个，不能存在相同地址。 该字段为null表示取消该子网NTP的设置，不能为””(空字符串)。 opt_name配置为“addresstime”，则该值表示是子网租约到期时间，取值格式有两种，取-1，表示无限租约；数字+h，数字范围是1~30000，比如5h。
	OptValue *string `json:"opt_value,omitempty"`
}

func (o ExtraDhcpOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExtraDhcpOption struct{}"
	}

	return strings.Join([]string{"ExtraDhcpOption", string(data)}, " ")
}

type ExtraDhcpOptionOptName struct {
	value string
}

type ExtraDhcpOptionOptNameEnum struct {
	NTP         ExtraDhcpOptionOptName
	ADDRESSTIME ExtraDhcpOptionOptName
}

func GetExtraDhcpOptionOptNameEnum() ExtraDhcpOptionOptNameEnum {
	return ExtraDhcpOptionOptNameEnum{
		NTP: ExtraDhcpOptionOptName{
			value: "ntp",
		},
		ADDRESSTIME: ExtraDhcpOptionOptName{
			value: "addresstime",
		},
	}
}

func (c ExtraDhcpOptionOptName) Value() string {
	return c.value
}

func (c ExtraDhcpOptionOptName) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ExtraDhcpOptionOptName) UnmarshalJSON(b []byte) error {
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
