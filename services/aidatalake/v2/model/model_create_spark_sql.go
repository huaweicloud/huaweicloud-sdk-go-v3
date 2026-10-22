package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateSparkSql **参数解释**：SparkSql引擎配置。
type CreateSparkSql struct {

	// **参数解释**：引擎类型。 **约束限制**：不涉及。 **取值范围**：SparkJob。 **默认取值**：不涉及。
	Type *CreateSparkSqlType `json:"type,omitempty"`

	// **参数解释**：SparkSql参数配置。 **约束限制**：最大数量为100。 **取值范围**：不涉及。 **默认取值**：不涉及。
	SparkConfig map[string]string `json:"spark_config,omitempty"`

	StagingConfig *StagingConfig `json:"staging_config"`

	ClusterConfig *SparkClusterConfig `json:"cluster_config"`

	// **参数解释**：端点最大并发SQL数量。 **约束限制**：不涉及。 **取值范围**：1~1000。 **默认取值**：300。
	MaxEndpointConcurrentQueries *int32 `json:"max_endpoint_concurrent_queries,omitempty"`

	Image *V2ImageBriefInfoAdmin `json:"image,omitempty"`
}

func (o CreateSparkSql) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSparkSql struct{}"
	}

	return strings.Join([]string{"CreateSparkSql", string(data)}, " ")
}

type CreateSparkSqlType struct {
	value string
}

type CreateSparkSqlTypeEnum struct {
	SPARK_SQL CreateSparkSqlType
}

func GetCreateSparkSqlTypeEnum() CreateSparkSqlTypeEnum {
	return CreateSparkSqlTypeEnum{
		SPARK_SQL: CreateSparkSqlType{
			value: "SparkSql",
		},
	}
}

func (c CreateSparkSqlType) Value() string {
	return c.value
}

func (c CreateSparkSqlType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateSparkSqlType) UnmarshalJSON(b []byte) error {
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
