package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 日志接入偏移时间
type AccessConfigTimeOffsetCreate struct {

	// 偏移时间。 当\"unit\"选择\"day\"时，范围为1~7天。 当\"unit\"选择\"hour\"时，范围为1~168小时。 当\"unit\"选择\"sec\"时，范围为1~604800秒。
	Offset int64 `json:"offset"`

	// 偏移时间单位。day ：天，hour：小时，sec：秒
	Unit AccessConfigTimeOffsetCreateUnit `json:"unit"`
}

func (o AccessConfigTimeOffsetCreate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AccessConfigTimeOffsetCreate struct{}"
	}

	return strings.Join([]string{"AccessConfigTimeOffsetCreate", string(data)}, " ")
}

type AccessConfigTimeOffsetCreateUnit struct {
	value string
}

type AccessConfigTimeOffsetCreateUnitEnum struct {
	DAY  AccessConfigTimeOffsetCreateUnit
	HOUR AccessConfigTimeOffsetCreateUnit
	SEC  AccessConfigTimeOffsetCreateUnit
}

func GetAccessConfigTimeOffsetCreateUnitEnum() AccessConfigTimeOffsetCreateUnitEnum {
	return AccessConfigTimeOffsetCreateUnitEnum{
		DAY: AccessConfigTimeOffsetCreateUnit{
			value: "day",
		},
		HOUR: AccessConfigTimeOffsetCreateUnit{
			value: "hour",
		},
		SEC: AccessConfigTimeOffsetCreateUnit{
			value: "sec",
		},
	}
}

func (c AccessConfigTimeOffsetCreateUnit) Value() string {
	return c.value
}

func (c AccessConfigTimeOffsetCreateUnit) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AccessConfigTimeOffsetCreateUnit) UnmarshalJSON(b []byte) error {
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
