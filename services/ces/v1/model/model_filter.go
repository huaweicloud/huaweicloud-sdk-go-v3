package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Filter **参数解释**： 数据聚合的方式。 **约束限制**： period为1（原始值）时，filter字段不生效，参数值默认为average。period大于1时filter才起作用。 **取值范围**： 支持 average、variance、min、max、sum。 - average： 平均值 - variance：方差 - min：最小值 - max：最大值 - sum：求和 **默认取值**： 不涉及。
type Filter struct {
	value string
}

type FilterEnum struct {
	AVERAGE  Filter
	VARIANCE Filter
	MIN      Filter
	MAX      Filter
	SUM      Filter
}

func GetFilterEnum() FilterEnum {
	return FilterEnum{
		AVERAGE: Filter{
			value: "average",
		},
		VARIANCE: Filter{
			value: "variance",
		},
		MIN: Filter{
			value: "min",
		},
		MAX: Filter{
			value: "max",
		},
		SUM: Filter{
			value: "sum",
		},
	}
}

func (c Filter) Value() string {
	return c.value
}

func (c Filter) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *Filter) UnmarshalJSON(b []byte) error {
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
