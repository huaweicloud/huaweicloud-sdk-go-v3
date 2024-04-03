package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type QueueProperty struct {

	// 返回属性值对应的key值 computeEngine.maxInstances, 队列能启动的最大spark driver数量 computeEngine.maxPrefetchInstance, 队列预先启动的最大spark driver数量 job.maxConcurrent,单个spark driver能同时运行的最大任务数量 multipleSc.support,是否支持设置多个spark driver
	Key QueuePropertyKey `json:"key"`

	Value string `json:"value"`
}

func (o QueueProperty) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueueProperty struct{}"
	}

	return strings.Join([]string{"QueueProperty", string(data)}, " ")
}

type QueuePropertyKey struct {
	value string
}

type QueuePropertyKeyEnum struct {
	COMPUTE_ENGINE_MAX_INSTANCES         QueuePropertyKey
	COMPUTE_ENGINE_MAX_PREFETCH_INSTANCE QueuePropertyKey
	JOB_MAX_CONCURRENT                   QueuePropertyKey
	MULTIPLE_SC_SUPPORT                  QueuePropertyKey
}

func GetQueuePropertyKeyEnum() QueuePropertyKeyEnum {
	return QueuePropertyKeyEnum{
		COMPUTE_ENGINE_MAX_INSTANCES: QueuePropertyKey{
			value: "computeEngine.maxInstances",
		},
		COMPUTE_ENGINE_MAX_PREFETCH_INSTANCE: QueuePropertyKey{
			value: "computeEngine.maxPrefetchInstance",
		},
		JOB_MAX_CONCURRENT: QueuePropertyKey{
			value: "job.maxConcurrent",
		},
		MULTIPLE_SC_SUPPORT: QueuePropertyKey{
			value: "multipleSc.support",
		},
	}
}

func (c QueuePropertyKey) Value() string {
	return c.value
}

func (c QueuePropertyKey) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *QueuePropertyKey) UnmarshalJSON(b []byte) error {
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
