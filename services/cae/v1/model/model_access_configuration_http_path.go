package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// AccessConfigurationHttpPath 匹配路径和域名信息。
type AccessConfigurationHttpPath struct {

	// 域名/不填，不填时表示使用IP。
	Hostname *string `json:"hostname,omitempty"`

	// URL路径。
	Path *string `json:"path,omitempty"`

	// URL路径匹配模式，支持前缀匹配、正则匹配、精准匹配。
	UrlMatchMode *AccessConfigurationHttpPathUrlMatchMode `json:"url_match_mode,omitempty"`
}

func (o AccessConfigurationHttpPath) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AccessConfigurationHttpPath struct{}"
	}

	return strings.Join([]string{"AccessConfigurationHttpPath", string(data)}, " ")
}

type AccessConfigurationHttpPathUrlMatchMode struct {
	value string
}

type AccessConfigurationHttpPathUrlMatchModeEnum struct {
	STARTS_WITH AccessConfigurationHttpPathUrlMatchMode
	REGEX       AccessConfigurationHttpPathUrlMatchMode
	EQUAL_TO    AccessConfigurationHttpPathUrlMatchMode
}

func GetAccessConfigurationHttpPathUrlMatchModeEnum() AccessConfigurationHttpPathUrlMatchModeEnum {
	return AccessConfigurationHttpPathUrlMatchModeEnum{
		STARTS_WITH: AccessConfigurationHttpPathUrlMatchMode{
			value: "STARTS_WITH",
		},
		REGEX: AccessConfigurationHttpPathUrlMatchMode{
			value: "REGEX",
		},
		EQUAL_TO: AccessConfigurationHttpPathUrlMatchMode{
			value: "EQUAL_TO",
		},
	}
}

func (c AccessConfigurationHttpPathUrlMatchMode) Value() string {
	return c.value
}

func (c AccessConfigurationHttpPathUrlMatchMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AccessConfigurationHttpPathUrlMatchMode) UnmarshalJSON(b []byte) error {
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
