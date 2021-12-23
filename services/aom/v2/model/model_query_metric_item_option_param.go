package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 参数项。
type QueryMetricItemOptionParam struct {
	// 指标维度列表。

	Dimensions *[]Dimension `json:"dimensions,omitempty"`
	// |  取值范围 名称长度为1~255个字符 指标名称。

	MetricName *string `json:"metricName,omitempty"`
	// 取值范围 PAAS.CONTAINER、PAAS.NODE、PAAS.SLA、PAAS.AGGR、CUSTOMMETRICS等 指标命名空间。 PAAS.CONTAINER：应用指标； PAAS.NODE：节点指标； PAAS.SLA：SLA指标； PAAS.AGGR：集群指标； CUSTOMMETRICS：自定义指标

	Namespace QueryMetricItemOptionParamNamespace `json:"namespace"`
}

func (o QueryMetricItemOptionParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryMetricItemOptionParam struct{}"
	}

	return strings.Join([]string{"QueryMetricItemOptionParam", string(data)}, " ")
}

type QueryMetricItemOptionParamNamespace struct {
	value string
}

type QueryMetricItemOptionParamNamespaceEnum struct {
	PAAS_CONTAINER QueryMetricItemOptionParamNamespace
	PAAS_NODE      QueryMetricItemOptionParamNamespace
	PAAS_SLA       QueryMetricItemOptionParamNamespace
	PAAS_AGGR      QueryMetricItemOptionParamNamespace
	CUSTOMMETRICS  QueryMetricItemOptionParamNamespace
}

func GetQueryMetricItemOptionParamNamespaceEnum() QueryMetricItemOptionParamNamespaceEnum {
	return QueryMetricItemOptionParamNamespaceEnum{
		PAAS_CONTAINER: QueryMetricItemOptionParamNamespace{
			value: "PAAS.CONTAINER",
		},
		PAAS_NODE: QueryMetricItemOptionParamNamespace{
			value: "PAAS.NODE",
		},
		PAAS_SLA: QueryMetricItemOptionParamNamespace{
			value: "PAAS.SLA",
		},
		PAAS_AGGR: QueryMetricItemOptionParamNamespace{
			value: "PAAS.AGGR",
		},
		CUSTOMMETRICS: QueryMetricItemOptionParamNamespace{
			value: "CUSTOMMETRICS",
		},
	}
}

func (c QueryMetricItemOptionParamNamespace) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *QueryMetricItemOptionParamNamespace) UnmarshalJSON(b []byte) error {
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
