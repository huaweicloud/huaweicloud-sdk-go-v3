package model

import (
	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

// ShowTableResponse Response Object
type ShowTableResponse struct {

	// catalog名字
	CatalogName *string `json:"catalog_name,omitempty"`

	// 数据库名字
	DatabaseName *string `json:"database_name,omitempty"`

	// 表名字
	TableName *string `json:"table_name,omitempty"`

	// 表创建时间
	CreateTime *sdktime.SdkTime `json:"create_time,omitempty"`

	// 最后访问时间
	LastAccessTime *sdktime.SdkTime `json:"last_access_time,omitempty"`

	// 表元数据最后一次修改时间
	UpdateTime *sdktime.SdkTime `json:"update_time,omitempty"`

	// 上一次做列级别的统计信息计算的时间
	LastAnalyzedTime *sdktime.SdkTime `json:"last_analyzed_time,omitempty"`

	// 用户信息
	Owner *string `json:"owner,omitempty"`

	// 用户类型
	OwnerType *ShowTableResponseOwnerType `json:"owner_type,omitempty"`

	// 参数
	Parameters map[string]string `json:"parameters,omitempty"`

	// 表分区列的列表
	PartitionKeys *[]Column `json:"partition_keys,omitempty"`

	// 表保留时间
	Retention *int32 `json:"retention,omitempty"`

	StorageDescriptor *StorageDescriptor `json:"storage_descriptor,omitempty"`

	// 表类型
	TableType *ShowTableResponseTableType `json:"table_type,omitempty"`

	// 表描述信息
	Comments *string `json:"comments,omitempty"`

	// 如果表是视图，则为视图的扩展文本；否则为 null
	ViewExpandedText *string `json:"view_expanded_text,omitempty"`

	// 如果表是视图，则为视图的原始文本；否则为 null
	ViewOriginalText *string `json:"view_original_text,omitempty"`
	HttpStatusCode   int     `json:"-"`
}

func (o ShowTableResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTableResponse struct{}"
	}

	return strings.Join([]string{"ShowTableResponse", string(data)}, " ")
}

type ShowTableResponseOwnerType struct {
	value string
}

type ShowTableResponseOwnerTypeEnum struct {
	USER  ShowTableResponseOwnerType
	GROUP ShowTableResponseOwnerType
	ROLE  ShowTableResponseOwnerType
}

func GetShowTableResponseOwnerTypeEnum() ShowTableResponseOwnerTypeEnum {
	return ShowTableResponseOwnerTypeEnum{
		USER: ShowTableResponseOwnerType{
			value: "USER",
		},
		GROUP: ShowTableResponseOwnerType{
			value: "GROUP",
		},
		ROLE: ShowTableResponseOwnerType{
			value: "ROLE",
		},
	}
}

func (c ShowTableResponseOwnerType) Value() string {
	return c.value
}

func (c ShowTableResponseOwnerType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowTableResponseOwnerType) UnmarshalJSON(b []byte) error {
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

type ShowTableResponseTableType struct {
	value string
}

type ShowTableResponseTableTypeEnum struct {
	MANAGED_TABLE     ShowTableResponseTableType
	EXTERNAL_TABLE    ShowTableResponseTableType
	VIRTUAL_VIEW      ShowTableResponseTableType
	MATERIALIZED_VIEW ShowTableResponseTableType
}

func GetShowTableResponseTableTypeEnum() ShowTableResponseTableTypeEnum {
	return ShowTableResponseTableTypeEnum{
		MANAGED_TABLE: ShowTableResponseTableType{
			value: "MANAGED_TABLE",
		},
		EXTERNAL_TABLE: ShowTableResponseTableType{
			value: "EXTERNAL_TABLE",
		},
		VIRTUAL_VIEW: ShowTableResponseTableType{
			value: "VIRTUAL_VIEW",
		},
		MATERIALIZED_VIEW: ShowTableResponseTableType{
			value: "MATERIALIZED_VIEW",
		},
	}
}

func (c ShowTableResponseTableType) Value() string {
	return c.value
}

func (c ShowTableResponseTableType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowTableResponseTableType) UnmarshalJSON(b []byte) error {
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
