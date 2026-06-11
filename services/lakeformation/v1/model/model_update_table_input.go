package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

// UpdateTableInput 表输入模型。
type UpdateTableInput struct {

	// 表名称。只能包含中文、字母、数字、下划线、中划线，且长度为1~256个字符。
	TableName *string `json:"table_name,omitempty"`

	// 表格式。支持{HIVE,ICEBERG,LANCE}，默认值为HIVE。
	TableFormat *UpdateTableInputTableFormat `json:"table_format,omitempty"`

	// 表类型：MANAGED_TABLE-内表、EXTERNAL_TABLE-外表、VIRTUAL_VIEW-视图、MATERIALIZED_VIEW-物化视图、DICTIONARY_TABLE字典表，LAKE_TABLE内表。
	TableType *UpdateTableInputTableType `json:"table_type,omitempty"`

	// 表所有者。只能包含字母、数字和下划线，且长度为1~49个字符。
	Owner *string `json:"owner,omitempty"`

	// 所有者类型：USER-用户、GROUP-组、ROLE-角色。
	OwnerType *UpdateTableInputOwnerType `json:"owner_type,omitempty"`

	// 表创建时间。
	CreateTime *sdktime.SdkTime `json:"create_time,omitempty"`

	// 最近一次访问时间。
	LastAccessTime *sdktime.SdkTime `json:"last_access_time,omitempty"`

	// 最近一次分析统计时间。
	LastAnalyzedTime *sdktime.SdkTime `json:"last_analyzed_time,omitempty"`

	// 分区列的信息。
	PartitionKeys *[]Column `json:"partition_keys,omitempty"`

	// 表保留时间。
	Retention *int32 `json:"retention,omitempty"`

	StorageDescriptor *StorageDescriptor `json:"storage_descriptor,omitempty"`

	// 表参数信息，每个键是一个键字符串，不少于 1 个字节或超过 255 个字节 每个值是一个 UTF-8 字符串，不超过 10000 个字节
	Parameters map[string]string `json:"parameters,omitempty"`

	// 表描述信息。由用户创建表时输入，最大长度为4000个字符。
	Comments *string `json:"comments,omitempty"`

	// 如果表是视图，则为视图的扩展文本。
	ViewExpandedText *string `json:"view_expanded_text,omitempty"`

	// 如果表是视图，则为视图的原始文本。
	ViewOriginalText *string `json:"view_original_text,omitempty"`

	// 是否忽略内表建表时对Obs路径的限制。
	IgnoreObsChecked *bool `json:"ignore_obs_checked,omitempty"`

	// 用户端表id，创建时指定，不可修改。
	ExternalTableId *string `json:"external_table_id,omitempty"`

	// 数据概况统计开关。默认状态为开，修改table开关状态后，还需检查所属database的开关状态。例如：table与所属database开关同时打开，则数据概况统计开启。否则关闭
	DataStatisticEnable *bool `json:"data_statistic_enable,omitempty"`

	// 版本ID，可在修改时传入，默认为最新版本
	VersionId *string `json:"version_id,omitempty"`
}

func (o UpdateTableInput) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTableInput struct{}"
	}

	return strings.Join([]string{"UpdateTableInput", string(data)}, " ")
}

type UpdateTableInputTableFormat struct {
	value string
}

type UpdateTableInputTableFormatEnum struct {
	HIVE    UpdateTableInputTableFormat
	ICEBERG UpdateTableInputTableFormat
	LANCE   UpdateTableInputTableFormat
}

func GetUpdateTableInputTableFormatEnum() UpdateTableInputTableFormatEnum {
	return UpdateTableInputTableFormatEnum{
		HIVE: UpdateTableInputTableFormat{
			value: "HIVE",
		},
		ICEBERG: UpdateTableInputTableFormat{
			value: "ICEBERG",
		},
		LANCE: UpdateTableInputTableFormat{
			value: "LANCE",
		},
	}
}

func (c UpdateTableInputTableFormat) Value() string {
	return c.value
}

func (c UpdateTableInputTableFormat) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateTableInputTableFormat) UnmarshalJSON(b []byte) error {
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

type UpdateTableInputTableType struct {
	value string
}

type UpdateTableInputTableTypeEnum struct {
	MANAGED_TABLE     UpdateTableInputTableType
	EXTERNAL_TABLE    UpdateTableInputTableType
	VIRTUAL_VIEW      UpdateTableInputTableType
	MATERIALIZED_VIEW UpdateTableInputTableType
	DICTIONARY_TABLE  UpdateTableInputTableType
	LAKE_TABLE        UpdateTableInputTableType
}

func GetUpdateTableInputTableTypeEnum() UpdateTableInputTableTypeEnum {
	return UpdateTableInputTableTypeEnum{
		MANAGED_TABLE: UpdateTableInputTableType{
			value: "MANAGED_TABLE",
		},
		EXTERNAL_TABLE: UpdateTableInputTableType{
			value: "EXTERNAL_TABLE",
		},
		VIRTUAL_VIEW: UpdateTableInputTableType{
			value: "VIRTUAL_VIEW",
		},
		MATERIALIZED_VIEW: UpdateTableInputTableType{
			value: "MATERIALIZED_VIEW",
		},
		DICTIONARY_TABLE: UpdateTableInputTableType{
			value: "DICTIONARY_TABLE",
		},
		LAKE_TABLE: UpdateTableInputTableType{
			value: "LAKE_TABLE",
		},
	}
}

func (c UpdateTableInputTableType) Value() string {
	return c.value
}

func (c UpdateTableInputTableType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateTableInputTableType) UnmarshalJSON(b []byte) error {
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

type UpdateTableInputOwnerType struct {
	value string
}

type UpdateTableInputOwnerTypeEnum struct {
	USER  UpdateTableInputOwnerType
	GROUP UpdateTableInputOwnerType
	ROLE  UpdateTableInputOwnerType
}

func GetUpdateTableInputOwnerTypeEnum() UpdateTableInputOwnerTypeEnum {
	return UpdateTableInputOwnerTypeEnum{
		USER: UpdateTableInputOwnerType{
			value: "USER",
		},
		GROUP: UpdateTableInputOwnerType{
			value: "GROUP",
		},
		ROLE: UpdateTableInputOwnerType{
			value: "ROLE",
		},
	}
}

func (c UpdateTableInputOwnerType) Value() string {
	return c.value
}

func (c UpdateTableInputOwnerType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateTableInputOwnerType) UnmarshalJSON(b []byte) error {
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
