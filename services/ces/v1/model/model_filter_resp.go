package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// FilterResp **参数解释**： 数据聚合的方式。 **取值范围**： 支持 average、variance、min、max、sum。 - average： 平均值 - variance：方差 - min：最小值 - max：最大值 - sum：求和
type FilterResp struct {
	value string
}

type FilterRespEnum struct {
	AVERAGE  FilterResp
	VARIANCE FilterResp
	MIN      FilterResp
	MAX      FilterResp
	SUM      FilterResp
}

func GetFilterRespEnum() FilterRespEnum {
	return FilterRespEnum{
		AVERAGE: FilterResp{
			value: "average",
		},
		VARIANCE: FilterResp{
			value: "variance",
		},
		MIN: FilterResp{
			value: "min",
		},
		MAX: FilterResp{
			value: "max",
		},
		SUM: FilterResp{
			value: "sum",
		},
	}
}

func (c FilterResp) Value() string {
	return c.value
}

func (c FilterResp) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *FilterResp) UnmarshalJSON(b []byte) error {
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
