package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// V2ImageCategory **参数解释**：自定义镜像分类。 **取值范围**：   - RAY_BASIC：基于Ray构建的自定义镜像。   - SPARK_JOB：基于SparkJob构建的自定义镜像。   - SPARK_SQL：基于SparkSql构建的自定义镜像。   - FLINK_JOB：基于FlinkJob构建的自定义镜像。   - AURA_JOB：基于AuraJobV2构建的自定义镜像。
type V2ImageCategory struct {
	value string
}

type V2ImageCategoryEnum struct {
	RAY_BASIC V2ImageCategory
	SPARK_JOB V2ImageCategory
	SPARK_SQL V2ImageCategory
	FLINK_JOB V2ImageCategory
	AURA_JOB  V2ImageCategory
}

func GetV2ImageCategoryEnum() V2ImageCategoryEnum {
	return V2ImageCategoryEnum{
		RAY_BASIC: V2ImageCategory{
			value: "RAY_BASIC",
		},
		SPARK_JOB: V2ImageCategory{
			value: "SPARK_JOB",
		},
		SPARK_SQL: V2ImageCategory{
			value: "SPARK_SQL",
		},
		FLINK_JOB: V2ImageCategory{
			value: "FLINK_JOB",
		},
		AURA_JOB: V2ImageCategory{
			value: "AURA_JOB",
		},
	}
}

func (c V2ImageCategory) Value() string {
	return c.value
}

func (c V2ImageCategory) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *V2ImageCategory) UnmarshalJSON(b []byte) error {
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
