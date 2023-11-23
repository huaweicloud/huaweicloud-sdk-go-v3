package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// SwitchIpGroupRequestBody 设置实例均衡负载IP的访问黑白名单请求体。
type SwitchIpGroupRequestBody struct {

	// 类型选项，取值： - whiteList：白名单，只允许指定ip或网段访问。 - blackList：黑名单，不允许指定ip或网段访问。
	Type SwitchIpGroupRequestBodyType `json:"type"`

	// true 开启，false 关闭。
	Enabled bool `json:"enabled"`

	// IP地址组中包含的IP或网段列表。
	IpGroups []SwitchIpGroupRequestBodyIpGroups `json:"ip_groups"`
}

func (o SwitchIpGroupRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SwitchIpGroupRequestBody struct{}"
	}

	return strings.Join([]string{"SwitchIpGroupRequestBody", string(data)}, " ")
}

type SwitchIpGroupRequestBodyType struct {
	value string
}

type SwitchIpGroupRequestBodyTypeEnum struct {
	WHITE_LIST SwitchIpGroupRequestBodyType
	BLACK_LIST SwitchIpGroupRequestBodyType
}

func GetSwitchIpGroupRequestBodyTypeEnum() SwitchIpGroupRequestBodyTypeEnum {
	return SwitchIpGroupRequestBodyTypeEnum{
		WHITE_LIST: SwitchIpGroupRequestBodyType{
			value: "whiteList",
		},
		BLACK_LIST: SwitchIpGroupRequestBodyType{
			value: "blackList",
		},
	}
}

func (c SwitchIpGroupRequestBodyType) Value() string {
	return c.value
}

func (c SwitchIpGroupRequestBodyType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SwitchIpGroupRequestBodyType) UnmarshalJSON(b []byte) error {
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
