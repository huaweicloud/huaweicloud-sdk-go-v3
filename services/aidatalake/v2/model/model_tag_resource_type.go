package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// TagResourceType **参数解释**：资源类型，用于指定需要打标签的资源类型，取值包括：   - aidatalake-workspace：工作空间。   - aidatalake-endpoint-ray：Ray引擎端点。   - aidatalake-endpoint-aura：Aura引擎端点。   - aidatalake-endpoint-spark：Spark引擎端点。   - aidatalake-endpoint-flink：Flink引擎端点。   - aidatalake-compute：计算资源。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
type TagResourceType struct {
	value string
}

type TagResourceTypeEnum struct {
	AIDATALAKE_WORKSPACE      TagResourceType
	AIDATALAKE_ENDPOINT_RAY   TagResourceType
	AIDATALAKE_ENDPOINT_AURA  TagResourceType
	AIDATALAKE_ENDPOINT_SPARK TagResourceType
	AIDATALAKE_ENDPOINT_FLINK TagResourceType
	AIDATALAKE_COMPUTE        TagResourceType
}

func GetTagResourceTypeEnum() TagResourceTypeEnum {
	return TagResourceTypeEnum{
		AIDATALAKE_WORKSPACE: TagResourceType{
			value: "aidatalake-workspace",
		},
		AIDATALAKE_ENDPOINT_RAY: TagResourceType{
			value: "aidatalake-endpoint-ray",
		},
		AIDATALAKE_ENDPOINT_AURA: TagResourceType{
			value: "aidatalake-endpoint-aura",
		},
		AIDATALAKE_ENDPOINT_SPARK: TagResourceType{
			value: "aidatalake-endpoint-spark",
		},
		AIDATALAKE_ENDPOINT_FLINK: TagResourceType{
			value: "aidatalake-endpoint-flink",
		},
		AIDATALAKE_COMPUTE: TagResourceType{
			value: "aidatalake-compute",
		},
	}
}

func (c TagResourceType) Value() string {
	return c.value
}

func (c TagResourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TagResourceType) UnmarshalJSON(b []byte) error {
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
