package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// SqlTypeConfigOption **参数解释**: SQL采集类型列表。默认取值[]，表示采集所有SQL语句。 **约束限制**: 不涉及。 **取值范围**: 不涉及。 **默认取值**: 不涉及。
type SqlTypeConfigOption struct {

	// **参数解释**: SQL类型的归类名称。 **约束限制**: - 对常用SQL类型，简单归类为6个类别，每个类别对应一组固定的采集SQL语句类型列表，采用前缀进行匹配。 - 对于其他场景，可以使用自定义类别，允许按需自定义采集SQL的语句前缀。  **取值范围**: - all：不区分SQL类型，全部采集。 - ddl：只采集DDL语句类别，包含如下SQL语句类型： create, alter, drop, truncate, reindex, vacuum, analyze, declare, move, close。 - dml：只采集DML语句类型，包含如下SQL语句类型： insert, update, delete, merge, show, explain, prepare, lock, copy, execute, deallocate。 - dcl：只采集DCL语句类型，包含如下SQL语句类型： grant, revoke, reassign, set。 - tcl：只采集TCL语句类型，包含如下SQL语句类型： begin, commit, rollback, start, savepoint, checkpoint, release savepoint。 - dql：只采集DQL语句类型，包含如下SQL语句类型： select。 - custom：采集自定义语句类型。需在prefixes字段中，填写要采集的SQL语句前缀片段。  **默认取值**: 不涉及。
	Category SqlTypeConfigOptionCategory `json:"category"`

	// **参数解释**: 针对custom自定义类别，指定要采集的SQL语句类型的列表，默认取值为[]。采集过程中，采用前缀方式对SQL文本进行匹配。 **约束限制**: - category取值custom时，此参数必填，不可为空。 - category取值其他类别时，此参数忽略。
	Prefixes *[]string `json:"prefixes,omitempty"`
}

func (o SqlTypeConfigOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SqlTypeConfigOption struct{}"
	}

	return strings.Join([]string{"SqlTypeConfigOption", string(data)}, " ")
}

type SqlTypeConfigOptionCategory struct {
	value string
}

type SqlTypeConfigOptionCategoryEnum struct {
	ALL    SqlTypeConfigOptionCategory
	DDL    SqlTypeConfigOptionCategory
	DML    SqlTypeConfigOptionCategory
	DCL    SqlTypeConfigOptionCategory
	TCL    SqlTypeConfigOptionCategory
	DQL    SqlTypeConfigOptionCategory
	CUSTOM SqlTypeConfigOptionCategory
}

func GetSqlTypeConfigOptionCategoryEnum() SqlTypeConfigOptionCategoryEnum {
	return SqlTypeConfigOptionCategoryEnum{
		ALL: SqlTypeConfigOptionCategory{
			value: "all",
		},
		DDL: SqlTypeConfigOptionCategory{
			value: "ddl",
		},
		DML: SqlTypeConfigOptionCategory{
			value: "dml",
		},
		DCL: SqlTypeConfigOptionCategory{
			value: "dcl",
		},
		TCL: SqlTypeConfigOptionCategory{
			value: "tcl",
		},
		DQL: SqlTypeConfigOptionCategory{
			value: "dql",
		},
		CUSTOM: SqlTypeConfigOptionCategory{
			value: "custom",
		},
	}
}

func (c SqlTypeConfigOptionCategory) Value() string {
	return c.value
}

func (c SqlTypeConfigOptionCategory) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SqlTypeConfigOptionCategory) UnmarshalJSON(b []byte) error {
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
