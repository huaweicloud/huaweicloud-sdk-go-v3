package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// RollupFilter 聚合规则，last:最新值,max:最大值,min:最小值,average:平均值,sum:求和值
type RollupFilter struct {
	value string
}

type RollupFilterEnum struct {
	LAST    RollupFilter
	MAX     RollupFilter
	MIN     RollupFilter
	AVERAGE RollupFilter
	SUM     RollupFilter
}

func GetRollupFilterEnum() RollupFilterEnum {
	return RollupFilterEnum{
		LAST: RollupFilter{
			value: "last",
		},
		MAX: RollupFilter{
			value: "max",
		},
		MIN: RollupFilter{
			value: "min",
		},
		AVERAGE: RollupFilter{
			value: "average",
		},
		SUM: RollupFilter{
			value: "sum",
		},
	}
}

func (c RollupFilter) Value() string {
	return c.value
}

func (c RollupFilter) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *RollupFilter) UnmarshalJSON(b []byte) error {
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
