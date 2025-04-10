package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// LbAccessControlSettings 负载均衡访问控制配置。
type LbAccessControlSettings struct {

	// true 开启，false 关闭。
	Enabled bool `json:"enabled"`

	// 黑白名单类型，blackList黑名单，whiteList白名单，仅支持设置黑名单或白名单中的一种，当配置切换时，原配置会失效。
	Type LbAccessControlSettingsType `json:"type"`

	// IP地址组中包含的IP或网段列表。
	IpGroups []IpGroupsDetail `json:"ip_groups"`
}

func (o LbAccessControlSettings) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LbAccessControlSettings struct{}"
	}

	return strings.Join([]string{"LbAccessControlSettings", string(data)}, " ")
}

type LbAccessControlSettingsType struct {
	value string
}

type LbAccessControlSettingsTypeEnum struct {
	BLACK_LIST LbAccessControlSettingsType
	WHITE_LIST LbAccessControlSettingsType
}

func GetLbAccessControlSettingsTypeEnum() LbAccessControlSettingsTypeEnum {
	return LbAccessControlSettingsTypeEnum{
		BLACK_LIST: LbAccessControlSettingsType{
			value: "blackList",
		},
		WHITE_LIST: LbAccessControlSettingsType{
			value: "whiteList",
		},
	}
}

func (c LbAccessControlSettingsType) Value() string {
	return c.value
}

func (c LbAccessControlSettingsType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *LbAccessControlSettingsType) UnmarshalJSON(b []byte) error {
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
