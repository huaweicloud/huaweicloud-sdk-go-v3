package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 日志接入偏移时间
type AccessConfigTimeOffset struct {

	// 偏移时间。 当\"unit\"选择\"day\"时，范围为1~7天。 当\"unit\"选择\"hour\"时，范围为1~168小时。 当\"unit\"选择\"sec\"时，范围为1~604800秒。
	Offset int64 `json:"offset"`

	// 偏移时间单位。day ：天，hour：小时，sec：秒
	Unit AccessConfigTimeOffsetUnit `json:"unit"`
}

func (o AccessConfigTimeOffset) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AccessConfigTimeOffset struct{}"
	}

	return strings.Join([]string{"AccessConfigTimeOffset", string(data)}, " ")
}

type AccessConfigTimeOffsetUnit struct {
	value string
}

type AccessConfigTimeOffsetUnitEnum struct {
	DAY  AccessConfigTimeOffsetUnit
	HOUR AccessConfigTimeOffsetUnit
	SEC  AccessConfigTimeOffsetUnit
}

func GetAccessConfigTimeOffsetUnitEnum() AccessConfigTimeOffsetUnitEnum {
	return AccessConfigTimeOffsetUnitEnum{
		DAY: AccessConfigTimeOffsetUnit{
			value: "day",
		},
		HOUR: AccessConfigTimeOffsetUnit{
			value: "hour",
		},
		SEC: AccessConfigTimeOffsetUnit{
			value: "sec",
		},
	}
}

func (c AccessConfigTimeOffsetUnit) Value() string {
	return c.value
}

func (c AccessConfigTimeOffsetUnit) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AccessConfigTimeOffsetUnit) UnmarshalJSON(b []byte) error {
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
