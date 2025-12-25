package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// TableCategory **参数解释**: 表类别 - STREAMING 流式传输 - INDEX 索引 - APPLICATION 应用 - TENANT_OBS 租户对象存储 - LAKE 数据湖  **约束限制** 不涉及 **取值范围**: - STREAMING - INDEX - APPLICATION - TENANT_OBS - LAKE  **默认值** 不涉及
type TableCategory struct {
	value string
}

type TableCategoryEnum struct {
	STREAMING   TableCategory
	INDEX       TableCategory
	APPLICATION TableCategory
	TENANT_OBS  TableCategory
	LAKE        TableCategory
}

func GetTableCategoryEnum() TableCategoryEnum {
	return TableCategoryEnum{
		STREAMING: TableCategory{
			value: "STREAMING",
		},
		INDEX: TableCategory{
			value: "INDEX",
		},
		APPLICATION: TableCategory{
			value: "APPLICATION",
		},
		TENANT_OBS: TableCategory{
			value: "TENANT_OBS",
		},
		LAKE: TableCategory{
			value: "LAKE",
		},
	}
}

func (c TableCategory) Value() string {
	return c.value
}

func (c TableCategory) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TableCategory) UnmarshalJSON(b []byte) error {
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
