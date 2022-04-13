package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

// match字段数据结构。
type Match struct {
	// 键，表示要匹配的字段。  key不能重复，value为匹配的值。  此字段为固定字典值。  不允许为空字符串。  当前key的参数值仅支持resource_name，此时value的参数值为专属主机名称。

	Key MatchKey `json:"key"`
	// 值。  每个值最大长度为255个unicode字符。  不可以为空 。

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
}

func GetMatchKeyEnum() MatchKeyEnum {
	return MatchKeyEnum{
		RESOURCE_NAME: MatchKey{
			value: "resource_name",
		},
	}
}

func (c MatchKey) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *MatchKey) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
