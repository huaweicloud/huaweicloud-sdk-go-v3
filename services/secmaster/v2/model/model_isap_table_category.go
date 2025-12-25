package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// IsapTableCategory **参数解释**: 目录类型 - STREAMING 实时流 - INDEX 索引 - APPLICATION 应用 - TENANT_BUCKET 租户桶 - LAKE 数据湖  **约束限制** 不涉及 **取值范围**: - STREAMING - INDEX - APPLICATION - TENANT_BUCKET - LAKE  **默认值** 不涉及
type IsapTableCategory struct {
	value string
}

type IsapTableCategoryEnum struct {
	STREAMING     IsapTableCategory
	INDEX         IsapTableCategory
	APPLICATION   IsapTableCategory
	TENANT_BUCKET IsapTableCategory
	LAKE          IsapTableCategory
}

func GetIsapTableCategoryEnum() IsapTableCategoryEnum {
	return IsapTableCategoryEnum{
		STREAMING: IsapTableCategory{
			value: "STREAMING",
		},
		INDEX: IsapTableCategory{
			value: "INDEX",
		},
		APPLICATION: IsapTableCategory{
			value: "APPLICATION",
		},
		TENANT_BUCKET: IsapTableCategory{
			value: "TENANT_BUCKET",
		},
		LAKE: IsapTableCategory{
			value: "LAKE",
		},
	}
}

func (c IsapTableCategory) Value() string {
	return c.value
}

func (c IsapTableCategory) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *IsapTableCategory) UnmarshalJSON(b []byte) error {
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
