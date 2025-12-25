package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// SkuEnum **参数解释**: sku类型 - FLOW_DATA_BANDWIDTH 数据带宽 - CSS_CAPACITY 索引集群容量 - PAIMON_CAPACITY 数据库容量 - OBS_CAPACITY 对象存储容量 - JOB_CAPACITY 作业容量 - AD_HOC_COUNT 即席查询数量  **约束限制** 不涉及 **取值范围**: - FLOW_DATA_BANDWIDTH - CSS_CAPACITY - PAIMON_CAPACITY - OBS_CAPACITY - JOB_CAPACITY - AD_HOC_COUNT  **默认值** 不涉及
type SkuEnum struct {
	value string
}

type SkuEnumEnum struct {
	FLOW_DATA_BANDWIDTH SkuEnum
	CSS_CAPACITY        SkuEnum
	PAIMON_CAPACITY     SkuEnum
	OBS_CAPACITY        SkuEnum
	JOB_CAPACITY        SkuEnum
	AD_HOC_COUNT        SkuEnum
}

func GetSkuEnumEnum() SkuEnumEnum {
	return SkuEnumEnum{
		FLOW_DATA_BANDWIDTH: SkuEnum{
			value: "FLOW_DATA_BANDWIDTH",
		},
		CSS_CAPACITY: SkuEnum{
			value: "CSS_CAPACITY",
		},
		PAIMON_CAPACITY: SkuEnum{
			value: "PAIMON_CAPACITY",
		},
		OBS_CAPACITY: SkuEnum{
			value: "OBS_CAPACITY",
		},
		JOB_CAPACITY: SkuEnum{
			value: "JOB_CAPACITY",
		},
		AD_HOC_COUNT: SkuEnum{
			value: "AD_HOC_COUNT",
		},
	}
}

func (c SkuEnum) Value() string {
	return c.value
}

func (c SkuEnum) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SkuEnum) UnmarshalJSON(b []byte) error {
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
