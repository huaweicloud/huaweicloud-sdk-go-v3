package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 搜索字段
type Match struct {

	// 键。  当前key的参数值只能取“resource_name”，此时value的参数值为资源名称。
	Key MatchKey `json:"key"`

	// 值。  当前key的参数值只能取“resource_name”，此时value的参数值为资源名称。
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

func (c MatchKey) Value() string {
	return c.value
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
