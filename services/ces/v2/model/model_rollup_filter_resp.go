package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// RollupFilterResp **参数解释** 聚合规则 **取值范围** - last:最新值 - max:最大值 - min:最小值 - average:平均值 - sum:求和值
type RollupFilterResp struct {
	value string
}

type RollupFilterRespEnum struct {
	LAST    RollupFilterResp
	MAX     RollupFilterResp
	MIN     RollupFilterResp
	AVERAGE RollupFilterResp
	SUM     RollupFilterResp
}

func GetRollupFilterRespEnum() RollupFilterRespEnum {
	return RollupFilterRespEnum{
		LAST: RollupFilterResp{
			value: "last",
		},
		MAX: RollupFilterResp{
			value: "max",
		},
		MIN: RollupFilterResp{
			value: "min",
		},
		AVERAGE: RollupFilterResp{
			value: "average",
		},
		SUM: RollupFilterResp{
			value: "sum",
		},
	}
}

func (c RollupFilterResp) Value() string {
	return c.value
}

func (c RollupFilterResp) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *RollupFilterResp) UnmarshalJSON(b []byte) error {
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
