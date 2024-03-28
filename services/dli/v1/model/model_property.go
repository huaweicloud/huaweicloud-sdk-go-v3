package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type Property struct {

	// 返回属性值对应的key值 computeEngine.maxInstances, 队列能启动的最大spark driver数量 computeEngine.maxPrefetchInstance, 队列预先启动的最大spark driver数量 job.maxConcurrent,单个spark driver能同时运行的最大任务数量 multipleSc.support,是否支持设置多个spark driver
	Key PropertyKey `json:"key"`

	Value string `json:"value"`
}

func (o Property) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Property struct{}"
	}

	return strings.Join([]string{"Property", string(data)}, " ")
}

type PropertyKey struct {
	value string
}

type PropertyKeyEnum struct {
	COMPUTE_ENGINE_MAX_INSTANCES         PropertyKey
	COMPUTE_ENGINE_MAX_PREFETCH_INSTANCE PropertyKey
	JOB_MAX_CONCURRENT                   PropertyKey
	MULTIPLE_SC_SUPPORT                  PropertyKey
}

func GetPropertyKeyEnum() PropertyKeyEnum {
	return PropertyKeyEnum{
		COMPUTE_ENGINE_MAX_INSTANCES: PropertyKey{
			value: "computeEngine.maxInstances",
		},
		COMPUTE_ENGINE_MAX_PREFETCH_INSTANCE: PropertyKey{
			value: "computeEngine.maxPrefetchInstance",
		},
		JOB_MAX_CONCURRENT: PropertyKey{
			value: "job.maxConcurrent",
		},
		MULTIPLE_SC_SUPPORT: PropertyKey{
			value: "multipleSc.support",
		},
	}
}

func (c PropertyKey) Value() string {
	return c.value
}

func (c PropertyKey) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PropertyKey) UnmarshalJSON(b []byte) error {
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
