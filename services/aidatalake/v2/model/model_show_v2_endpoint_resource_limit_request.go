package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowV2EndpointResourceLimitRequest Request Object
type ShowV2EndpointResourceLimitRequest struct {

	// **参数解释**：工作空间ID，可通过[获取工作空间ID](aidatalake_02_0044.xml)获取。 **约束限制**：不涉及。 **取值范围**：长度为32~36的英文字符、数字和中划线的组合。 **默认取值**：不涉及。
	WorkspaceId string `json:"workspace_id"`

	// **参数解释**：端点引擎类型。 **约束限制**：不涉及。 **取值范围**：   - AuraJob：AuraJob引擎。   - SparkSql：SparkSql引擎。   - SparkJob：SparkJob引擎。   - FlinkJob：FlinkJob引擎。   - AuraJobV2作业管理：AuraJobV2引擎。   - RayJob：RayJob引擎。   - RayCluster：RayCluster引擎。   - RayService：RayService引擎。 **默认取值**：不涉及。
	EngineType ShowV2EndpointResourceLimitRequestEngineType `json:"engine_type"`

	// **参数解释**：端点ID，可通过[查询端点列表](ListV2Endpoints.xml)获取端点ID。 **约束限制**：不涉及。 **取值范围**：长度为1~36的英文字符、数字和中划线的组合。 **默认取值**：不涉及。
	Id *string `json:"id,omitempty"`

	// **参数解释**：计算资源ID，可通过[查询计算资源列表](ListV2Computes.xml)获取计算资源ID。 **约束限制**：不涉及。 **取值范围**：长度为1-36个英文字母、数字和中划线的组合。 **默认取值**：不涉及。
	ComputeId string `json:"compute_id"`
}

func (o ShowV2EndpointResourceLimitRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowV2EndpointResourceLimitRequest struct{}"
	}

	return strings.Join([]string{"ShowV2EndpointResourceLimitRequest", string(data)}, " ")
}

type ShowV2EndpointResourceLimitRequestEngineType struct {
	value string
}

type ShowV2EndpointResourceLimitRequestEngineTypeEnum struct {
	AURA_JOB    ShowV2EndpointResourceLimitRequestEngineType
	SPARK_SQL   ShowV2EndpointResourceLimitRequestEngineType
	SPARK_JOB   ShowV2EndpointResourceLimitRequestEngineType
	FLINK_JOB   ShowV2EndpointResourceLimitRequestEngineType
	AURA_JOB_V2 ShowV2EndpointResourceLimitRequestEngineType
	RAY_JOB     ShowV2EndpointResourceLimitRequestEngineType
	RAY_CLUSTER ShowV2EndpointResourceLimitRequestEngineType
	RAY_SERVICE ShowV2EndpointResourceLimitRequestEngineType
}

func GetShowV2EndpointResourceLimitRequestEngineTypeEnum() ShowV2EndpointResourceLimitRequestEngineTypeEnum {
	return ShowV2EndpointResourceLimitRequestEngineTypeEnum{
		AURA_JOB: ShowV2EndpointResourceLimitRequestEngineType{
			value: "AuraJob",
		},
		SPARK_SQL: ShowV2EndpointResourceLimitRequestEngineType{
			value: "SparkSql",
		},
		SPARK_JOB: ShowV2EndpointResourceLimitRequestEngineType{
			value: "SparkJob",
		},
		FLINK_JOB: ShowV2EndpointResourceLimitRequestEngineType{
			value: "FlinkJob",
		},
		AURA_JOB_V2: ShowV2EndpointResourceLimitRequestEngineType{
			value: "AuraJobV2作业管理",
		},
		RAY_JOB: ShowV2EndpointResourceLimitRequestEngineType{
			value: "RayJob",
		},
		RAY_CLUSTER: ShowV2EndpointResourceLimitRequestEngineType{
			value: "RayCluster",
		},
		RAY_SERVICE: ShowV2EndpointResourceLimitRequestEngineType{
			value: "RayService",
		},
	}
}

func (c ShowV2EndpointResourceLimitRequestEngineType) Value() string {
	return c.value
}

func (c ShowV2EndpointResourceLimitRequestEngineType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowV2EndpointResourceLimitRequestEngineType) UnmarshalJSON(b []byte) error {
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
