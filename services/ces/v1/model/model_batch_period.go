package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// BatchPeriod 指标监控数据的聚合粒度，取值范围：1，300，1200，3600，14400，86400；1为监控资源的实时数据；300为聚合5分钟粒度数据，表示5分钟一个数据点；1200为聚合20分钟粒度数据，表示20分钟一个数据点；3600为聚合1小时粒度数据，表示1小时一个数据点；14400为聚合4小时粒度数据，表示4小时一个数据点；86400为聚合1天粒度数据，表示1天一个数据点；聚合解释可查看：“[聚合含义](https://support.huaweicloud.com/ces_faq/ces_faq_0009.html)”。
type BatchPeriod struct {
	value string
}

type BatchPeriodEnum struct {
	E_1     BatchPeriod
	E_60    BatchPeriod
	E_300   BatchPeriod
	E_1200  BatchPeriod
	E_3600  BatchPeriod
	E_14400 BatchPeriod
	E_86400 BatchPeriod
}

func GetBatchPeriodEnum() BatchPeriodEnum {
	return BatchPeriodEnum{
		E_1: BatchPeriod{
			value: "1",
		},
		E_60: BatchPeriod{
			value: "60",
		},
		E_300: BatchPeriod{
			value: "300",
		},
		E_1200: BatchPeriod{
			value: "1200",
		},
		E_3600: BatchPeriod{
			value: "3600",
		},
		E_14400: BatchPeriod{
			value: "14400",
		},
		E_86400: BatchPeriod{
			value: "86400",
		},
	}
}

func (c BatchPeriod) Value() string {
	return c.value
}

func (c BatchPeriod) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchPeriod) UnmarshalJSON(b []byte) error {
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
