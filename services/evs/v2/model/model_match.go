package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type Match struct {

	// 键。取值范围如下： resource_name：资源名称。 service_type：服务类型。
	Key MatchKey `json:"key"`

	// 值。最大长度255个字符。 key为“resource_name”时，value为模糊匹配。
	Value string `json:"value"`
}

func (o Match) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Match struct{}"
	}

	return strings.Join([]string{"Match", string(data)}, " ")
}

type MatchKey struct {
	value string
}

type MatchKeyEnum struct {
	RESOURCE_NAME MatchKey
	SERVICE_TYPE  MatchKey
}

func GetMatchKeyEnum() MatchKeyEnum {
	return MatchKeyEnum{
		RESOURCE_NAME: MatchKey{
			value: "resource_name",
		},
		SERVICE_TYPE: MatchKey{
			value: "service_type",
		},
	}
}

func (c MatchKey) Value() string {
	return c.value
}

func (c MatchKey) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *MatchKey) UnmarshalJSON(b []byte) error {
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
