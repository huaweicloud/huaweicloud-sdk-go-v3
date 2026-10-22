package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// V2EndpointUpdateEngineConfig **参数解释**：创建端点引擎配置。
type V2EndpointUpdateEngineConfig struct {

	// **参数解释**：引擎类型。 **约束限制**：不涉及。 **取值范围**：RayService。 **默认取值**：不涉及。
	Type *V2EndpointUpdateEngineConfigType `json:"type,omitempty"`

	// **参数解释**：Spark 参数配置。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	SparkConfig map[string]string `json:"spark_config,omitempty"`

	SchedulingConfig *UpdateSparkSchedulingConfig `json:"scheduling_config,omitempty"`

	Image *V2ImageBriefInfoAdmin `json:"image,omitempty"`

	ClusterConfig *RayClusterConfigV3 `json:"cluster_config,omitempty"`

	// **参数解释**：端点最大并发SQL数量。 **约束限制**：不涉及。 **取值范围**：1~1000。 **默认取值**：300。
	MaxEndpointConcurrentQueries *int32 `json:"max_endpoint_concurrent_queries,omitempty"`

	HistoryServerConfig *HistoryServerConfig `json:"history_server_config,omitempty"`

	RedisConfig *RedisConfig `json:"redis_config,omitempty"`

	StoragesConfig *StoragesConfig `json:"storages_config,omitempty"`

	// **参数解释**：Flink参数配置，Map的key为Flink配置参数名，value为对应的参数值。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	FlinkConfig map[string]string `json:"flink_config,omitempty"`

	// **参数解释**：工作组补充配置，和外部capacity对应。 **约束限制**：最多10个。
	WorkerGroupConfig *[]WorkerGroupConfig `json:"worker_group_config,omitempty"`

	// **参数解释**：开启后，通算工作组作业能够使用智算工作组空闲CPU及内存资源，该操作可能会导致智算任务因资源不足而排队。 **约束限制**：不涉及。 **取值范围**：   - true：开启通算工作组作业使用智算工作组空闲CPU及内存资源。   - false：关闭通算工作组作业使用智算工作组空闲CPU及内存资源。 **默认取值**：false。
	ExpandCpuResource *bool `json:"expand_cpu_resource,omitempty"`

	ServeConfig *ServeConfig `json:"serve_config,omitempty"`

	// **参数解释**：模型信息列表（可选）。更新时如果模型没有变化，可以不填此字段；如果需要更新模型，则填写新的模型信息。 **约束限制**：不涉及。
	DataInfos *[]DataInfo `json:"data_infos,omitempty"`
}

func (o V2EndpointUpdateEngineConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "V2EndpointUpdateEngineConfig struct{}"
	}

	return strings.Join([]string{"V2EndpointUpdateEngineConfig", string(data)}, " ")
}

type V2EndpointUpdateEngineConfigType struct {
	value string
}

type V2EndpointUpdateEngineConfigTypeEnum struct {
	RAY_SERVICE V2EndpointUpdateEngineConfigType
}

func GetV2EndpointUpdateEngineConfigTypeEnum() V2EndpointUpdateEngineConfigTypeEnum {
	return V2EndpointUpdateEngineConfigTypeEnum{
		RAY_SERVICE: V2EndpointUpdateEngineConfigType{
			value: "RayService",
		},
	}
}

func (c V2EndpointUpdateEngineConfigType) Value() string {
	return c.value
}

func (c V2EndpointUpdateEngineConfigType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *V2EndpointUpdateEngineConfigType) UnmarshalJSON(b []byte) error {
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
