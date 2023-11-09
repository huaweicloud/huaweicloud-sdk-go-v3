package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ListQueuePropertyRespProperties struct {

	// 返回属性值对应的key值 computeEngine.maxInstances, 队列能启动的最大spark driver数量 computeEngine.maxPrefetchInstance, 队列预先启动的最大spark driver数量 job.maxConcurrent,单个spark driver能同时运行的最大任务数量 multipleSc.support,是否支持设置多个spark driver
	Key ListQueuePropertyRespPropertiesKey `json:"key"`

	Value string `json:"value"`
}

func (o ListQueuePropertyRespProperties) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListQueuePropertyRespProperties struct{}"
	}

	return strings.Join([]string{"ListQueuePropertyRespProperties", string(data)}, " ")
}

type ListQueuePropertyRespPropertiesKey struct {
	value string
}

type ListQueuePropertyRespPropertiesKeyEnum struct {
	COMPUTE_ENGINE_MAX_INSTANCES         ListQueuePropertyRespPropertiesKey
	COMPUTE_ENGINE_MAX_PREFETCH_INSTANCE ListQueuePropertyRespPropertiesKey
	JOB_MAX_CONCURRENT                   ListQueuePropertyRespPropertiesKey
	MULTIPLE_SC_SUPPORT                  ListQueuePropertyRespPropertiesKey
}

func GetListQueuePropertyRespPropertiesKeyEnum() ListQueuePropertyRespPropertiesKeyEnum {
	return ListQueuePropertyRespPropertiesKeyEnum{
		COMPUTE_ENGINE_MAX_INSTANCES: ListQueuePropertyRespPropertiesKey{
			value: "computeEngine.maxInstances",
		},
		COMPUTE_ENGINE_MAX_PREFETCH_INSTANCE: ListQueuePropertyRespPropertiesKey{
			value: "computeEngine.maxPrefetchInstance",
		},
		JOB_MAX_CONCURRENT: ListQueuePropertyRespPropertiesKey{
			value: "job.maxConcurrent",
		},
		MULTIPLE_SC_SUPPORT: ListQueuePropertyRespPropertiesKey{
			value: "multipleSc.support",
		},
	}
}

func (c ListQueuePropertyRespPropertiesKey) Value() string {
	return c.value
}

func (c ListQueuePropertyRespPropertiesKey) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListQueuePropertyRespPropertiesKey) UnmarshalJSON(b []byte) error {
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
