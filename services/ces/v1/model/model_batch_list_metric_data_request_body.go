package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// BatchListMetricDataRequestBody
type BatchListMetricDataRequestBody struct {

	// 指标数据。数组长度最大500
	Metrics []MetricInfo `json:"metrics"`

	// 指标监控数据的聚合粒度，取值范围：1，300，1200，3600，14400，86400；1为监控资源的实时数据；300为聚合5分钟粒度数据，表示5分钟一个数据点；1200为聚合20分钟粒度数据，表示20分钟一个数据点；3600为聚合1小时粒度数据，表示1小时一个数据点；14400为聚合4小时粒度数据，表示4小时一个数据点；86400为聚合1天粒度数据，表示1天一个数据点；聚合解释可查看：“[聚合含义](https://support.huaweicloud.com/ces_faq/ces_faq_0009.html)”。
	Period BatchListMetricDataRequestBodyPeriod `json:"period"`

	Filter *Filter `json:"filter"`

	From int64 `json:"from"`

	To int64 `json:"to"`
}

func (o BatchListMetricDataRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchListMetricDataRequestBody struct{}"
	}

	return strings.Join([]string{"BatchListMetricDataRequestBody", string(data)}, " ")
}

type BatchListMetricDataRequestBodyPeriod struct {
	value string
}

type BatchListMetricDataRequestBodyPeriodEnum struct {
	E_1     BatchListMetricDataRequestBodyPeriod
	E_60    BatchListMetricDataRequestBodyPeriod
	E_300   BatchListMetricDataRequestBodyPeriod
	E_1200  BatchListMetricDataRequestBodyPeriod
	E_3600  BatchListMetricDataRequestBodyPeriod
	E_14400 BatchListMetricDataRequestBodyPeriod
	E_86400 BatchListMetricDataRequestBodyPeriod
}

func GetBatchListMetricDataRequestBodyPeriodEnum() BatchListMetricDataRequestBodyPeriodEnum {
	return BatchListMetricDataRequestBodyPeriodEnum{
		E_1: BatchListMetricDataRequestBodyPeriod{
			value: "1",
		},
		E_60: BatchListMetricDataRequestBodyPeriod{
			value: "60",
		},
		E_300: BatchListMetricDataRequestBodyPeriod{
			value: "300",
		},
		E_1200: BatchListMetricDataRequestBodyPeriod{
			value: "1200",
		},
		E_3600: BatchListMetricDataRequestBodyPeriod{
			value: "3600",
		},
		E_14400: BatchListMetricDataRequestBodyPeriod{
			value: "14400",
		},
		E_86400: BatchListMetricDataRequestBodyPeriod{
			value: "86400",
		},
	}
}

func (c BatchListMetricDataRequestBodyPeriod) Value() string {
	return c.value
}

func (c BatchListMetricDataRequestBodyPeriod) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchListMetricDataRequestBodyPeriod) UnmarshalJSON(b []byte) error {
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
