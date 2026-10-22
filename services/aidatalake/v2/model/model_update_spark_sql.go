package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type UpdateSparkSql struct {

	// **参数解释**：引擎类型。 **约束限制**：不涉及。 **取值范围**：SparkSql。 **默认取值**：不涉及。
	Type *UpdateSparkSqlType `json:"type,omitempty"`

	Image *V2ImageBriefInfoAdmin `json:"image,omitempty"`

	// **参数解释**：Spark 参数配置。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	SparkConfig map[string]string `json:"spark_config,omitempty"`

	ClusterConfig *UpdateSparkClusterConfig `json:"cluster_config,omitempty"`

	// **参数解释**：端点最大并发SQL数量。 **约束限制**：不涉及。 **取值范围**：1~1000。 **默认取值**：300。
	MaxEndpointConcurrentQueries *int32 `json:"max_endpoint_concurrent_queries,omitempty"`
}

func (o UpdateSparkSql) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSparkSql struct{}"
	}

	return strings.Join([]string{"UpdateSparkSql", string(data)}, " ")
}

type UpdateSparkSqlType struct {
	value string
}

type UpdateSparkSqlTypeEnum struct {
	SPARK_SQL UpdateSparkSqlType
}

func GetUpdateSparkSqlTypeEnum() UpdateSparkSqlTypeEnum {
	return UpdateSparkSqlTypeEnum{
		SPARK_SQL: UpdateSparkSqlType{
			value: "SparkSql",
		},
	}
}

func (c UpdateSparkSqlType) Value() string {
	return c.value
}

func (c UpdateSparkSqlType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateSparkSqlType) UnmarshalJSON(b []byte) error {
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
