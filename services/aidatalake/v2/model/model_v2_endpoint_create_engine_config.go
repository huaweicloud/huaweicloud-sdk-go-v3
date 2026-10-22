package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// V2EndpointCreateEngineConfig **参数解释**：创建端点引擎配置。
type V2EndpointCreateEngineConfig struct {

	// **参数解释**：引擎类型。 **约束限制**：不涉及。 **取值范围**：RayService。 **默认取值**：不涉及。
	Type *V2EndpointCreateEngineConfigType `json:"type,omitempty"`

	// **参数解释**：SparkSql参数配置。 **约束限制**：最大数量为100。 **取值范围**：不涉及。 **默认取值**：不涉及。
	SparkConfig map[string]string `json:"spark_config,omitempty"`

	SchedulingConfig *SparkSchedulingConfig `json:"scheduling_config,omitempty"`

	StagingConfig *AuraV2StagingConfig `json:"staging_config"`

	Image *V2ImageBriefInfoAdmin `json:"image"`

	ClusterConfig *RayClusterConfigV3 `json:"cluster_config"`

	// **参数解释**：端点最大并发SQL数量。 **约束限制**：不涉及。 **取值范围**：1~1000。 **默认取值**：300。
	MaxEndpointConcurrentQueries *int32 `json:"max_endpoint_concurrent_queries,omitempty"`

	HistoryServerConfig *HistoryServerConfig `json:"history_server_config,omitempty"`

	RedisConfig *RedisConfig `json:"redis_config,omitempty"`

	StoragesConfig *CreateStoragesConfig `json:"storages_config"`

	// **参数解释**：Flink参数配置，Map的key为Flink配置参数名，value为对应的参数值。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	FlinkConfig map[string]string `json:"flink_config,omitempty"`

	// **参数解释**：开启后，通算工作组作业能够使用智算工作组空闲CPU及内存资源，该操作可能会导致智算任务因资源不足而排队。 **约束限制**：不涉及。 **取值范围**：   - true：开启通算工作组作业使用智算工作组空闲CPU及内存资源。   - false：关闭通算工作组作业使用智算工作组空闲CPU及内存资源。 **默认取值**：false。
	ExpandCpuResource *bool `json:"expand_cpu_resource,omitempty"`

	// **参数解释**：工作组补充配置，和外部capacity对应。 **约束限制**：最多10个。
	WorkerGroupConfig []WorkerGroupConfig `json:"worker_group_config"`

	ServeConfig *ServeConfig `json:"serve_config"`

	// **参数解释**：模型信息列表（可选）。模型可以打包在镜像中（通过 working_dir 配置项指定），也可以通过此字段上传模型地址。创建时如果模型已预置在镜像中，可省略此字段；更新时如果模型没有变化，也可以不填。 **约束限制**：不涉及。
	DataInfos *[]DataInfo `json:"data_infos,omitempty"`
}

func (o V2EndpointCreateEngineConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "V2EndpointCreateEngineConfig struct{}"
	}

	return strings.Join([]string{"V2EndpointCreateEngineConfig", string(data)}, " ")
}

type V2EndpointCreateEngineConfigType struct {
	value string
}

type V2EndpointCreateEngineConfigTypeEnum struct {
	RAY_SERVICE V2EndpointCreateEngineConfigType
}

func GetV2EndpointCreateEngineConfigTypeEnum() V2EndpointCreateEngineConfigTypeEnum {
	return V2EndpointCreateEngineConfigTypeEnum{
		RAY_SERVICE: V2EndpointCreateEngineConfigType{
			value: "RayService",
		},
	}
}

func (c V2EndpointCreateEngineConfigType) Value() string {
	return c.value
}

func (c V2EndpointCreateEngineConfigType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *V2EndpointCreateEngineConfigType) UnmarshalJSON(b []byte) error {
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
