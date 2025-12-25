package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type CreateTableRequestBody struct {

	// 表名称
	TableName string `json:"table_name"`

	// 表别名
	TableAlias string `json:"table_alias"`

	// 表别名
	TableAliasEn *string `json:"table_alias_en,omitempty"`

	// 表别名
	TableAliasFr *string `json:"table_alias_fr,omitempty"`

	// **参数解释**: 目录类型 - STREAMING 实时流 - INDEX 索引 - APPLICATION 应用 - TENANT_BUCKET 租户桶 - LAKE 数据湖  **约束限制** 不涉及 **取值范围**: - STREAMING - INDEX - APPLICATION - TENANT_BUCKET - LAKE  **默认值** 不涉及
	Category CreateTableRequestBodyCategory `json:"category"`

	// **参数解释**: 表格式 - JSON Json格式 - DEBEZIUM_JSON Debezium JSON 格式 - CSV CSV格式 - PARQUET PARQUET格式 - ORC ORC格式  **约束限制** 不涉及 **取值范围**: - JSON - DEBEZIUM_JSON - CSV - PARQUET - ORC  **默认值** 不涉及
	Format CreateTableRequestBodyFormat `json:"format"`

	// 表相关描述
	Description *string `json:"description,omitempty"`

	// 表相关描述
	DescriptionEn *string `json:"description_en,omitempty"`

	// 表相关描述
	DescriptionFr *string `json:"description_fr,omitempty"`

	// 目录分组
	Directory *string `json:"directory,omitempty"`

	// 目录分组
	DirectoryEn *string `json:"directory_en,omitempty"`

	// 目录分组
	DirectoryFr *string `json:"directory_fr,omitempty"`

	DataClassification *DataClassification `json:"data_classification,omitempty"`

	Schema *IsapTableSchema `json:"schema"`

	StorageSetting *TableStorageSetting `json:"storage_setting"`

	DisplaySetting *TableDisplaySetting `json:"display_setting"`

	// **参数解释**: 创建政策 - SYS_INIT_INDEX_APP_TBL 系统初始化索引应用表 - DEFAULT 默认  **约束限制** 不涉及 **取值范围**: - SYS_INIT_INDEX_APP_TBL - DEFAULT  **默认值** 不涉及
	CreatePolicy *CreateTableRequestBodyCreatePolicy `json:"create_policy,omitempty"`
}

func (o CreateTableRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTableRequestBody struct{}"
	}

	return strings.Join([]string{"CreateTableRequestBody", string(data)}, " ")
}

type CreateTableRequestBodyCategory struct {
	value string
}

type CreateTableRequestBodyCategoryEnum struct {
	STREAMING     CreateTableRequestBodyCategory
	INDEX         CreateTableRequestBodyCategory
	APPLICATION   CreateTableRequestBodyCategory
	TENANT_BUCKET CreateTableRequestBodyCategory
	LAKE          CreateTableRequestBodyCategory
}

func GetCreateTableRequestBodyCategoryEnum() CreateTableRequestBodyCategoryEnum {
	return CreateTableRequestBodyCategoryEnum{
		STREAMING: CreateTableRequestBodyCategory{
			value: "STREAMING",
		},
		INDEX: CreateTableRequestBodyCategory{
			value: "INDEX",
		},
		APPLICATION: CreateTableRequestBodyCategory{
			value: "APPLICATION",
		},
		TENANT_BUCKET: CreateTableRequestBodyCategory{
			value: "TENANT_BUCKET",
		},
		LAKE: CreateTableRequestBodyCategory{
			value: "LAKE",
		},
	}
}

func (c CreateTableRequestBodyCategory) Value() string {
	return c.value
}

func (c CreateTableRequestBodyCategory) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateTableRequestBodyCategory) UnmarshalJSON(b []byte) error {
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

type CreateTableRequestBodyFormat struct {
	value string
}

type CreateTableRequestBodyFormatEnum struct {
	JSON          CreateTableRequestBodyFormat
	DEBEZIUM_JSON CreateTableRequestBodyFormat
	CSV           CreateTableRequestBodyFormat
	PARQUET       CreateTableRequestBodyFormat
	ORC           CreateTableRequestBodyFormat
}

func GetCreateTableRequestBodyFormatEnum() CreateTableRequestBodyFormatEnum {
	return CreateTableRequestBodyFormatEnum{
		JSON: CreateTableRequestBodyFormat{
			value: "JSON",
		},
		DEBEZIUM_JSON: CreateTableRequestBodyFormat{
			value: "DEBEZIUM_JSON",
		},
		CSV: CreateTableRequestBodyFormat{
			value: "CSV",
		},
		PARQUET: CreateTableRequestBodyFormat{
			value: "PARQUET",
		},
		ORC: CreateTableRequestBodyFormat{
			value: "ORC",
		},
	}
}

func (c CreateTableRequestBodyFormat) Value() string {
	return c.value
}

func (c CreateTableRequestBodyFormat) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateTableRequestBodyFormat) UnmarshalJSON(b []byte) error {
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

type CreateTableRequestBodyCreatePolicy struct {
	value string
}

type CreateTableRequestBodyCreatePolicyEnum struct {
	SYS_INIT_INDEX_APP_TBL CreateTableRequestBodyCreatePolicy
	DEFAULT                CreateTableRequestBodyCreatePolicy
}

func GetCreateTableRequestBodyCreatePolicyEnum() CreateTableRequestBodyCreatePolicyEnum {
	return CreateTableRequestBodyCreatePolicyEnum{
		SYS_INIT_INDEX_APP_TBL: CreateTableRequestBodyCreatePolicy{
			value: "SYS_INIT_INDEX_APP_TBL",
		},
		DEFAULT: CreateTableRequestBodyCreatePolicy{
			value: "DEFAULT",
		},
	}
}

func (c CreateTableRequestBodyCreatePolicy) Value() string {
	return c.value
}

func (c CreateTableRequestBodyCreatePolicy) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateTableRequestBodyCreatePolicy) UnmarshalJSON(b []byte) error {
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
