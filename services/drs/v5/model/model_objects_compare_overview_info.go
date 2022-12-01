package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 对象对比结果信息体。
type ObjectsCompareOverviewInfo struct {

	// 对象类型。取值： - DB：数据库。 - TABLE：表。 - VIEW：视图。 - EVENT：事件。 - ROUTINE：存储过程和函数。 - INDEX：索引。 - TRIGGER：触发器。 - SYNONYM：同义词。 - FUNCTION：函数。 - PROCEDURE：存储过程。 - TYPE：自定义类型。 - RULE：规则。 - DEFAULT_TYPE：缺省值。 - PLAN_GUIDE：执行计划。 - CONSTRAINT：约束。 - FILE_GROUP：文件组。 - PARTITION_FUNCTION：分区函数。 - PARTITION_SCHEME：分区方案。 - TABLE_COLLATION：表的排序规则。 - EXTENSIONS：插件。
	Type *ObjectsCompareOverviewInfoType `json:"type,omitempty"`

	// 源数量。
	SourceCount *int64 `json:"source_count,omitempty"`

	// 目标数量。
	TargetCount *int64 `json:"target_count,omitempty"`

	// 对比结果。取值： CONSISTENT：一致。 INCONSISTENT：不一致。 COMPARING：正在对比。 WAITING_FOR_COMPARISON：等待对比。 FAILED_TO_COMPARE：对比失败。 TARGET_DB_NOT_EXIST：目标库不存在。 CAN_NOT_COMPARE：无法对比。
	Status *ObjectsCompareOverviewInfoStatus `json:"status,omitempty"`
}

func (o ObjectsCompareOverviewInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ObjectsCompareOverviewInfo struct{}"
	}

	return strings.Join([]string{"ObjectsCompareOverviewInfo", string(data)}, " ")
}

type ObjectsCompareOverviewInfoType struct {
	value string
}

type ObjectsCompareOverviewInfoTypeEnum struct {
	DB                 ObjectsCompareOverviewInfoType
	TABLE              ObjectsCompareOverviewInfoType
	VIEW               ObjectsCompareOverviewInfoType
	EVENT              ObjectsCompareOverviewInfoType
	ROUTINE            ObjectsCompareOverviewInfoType
	INDEX              ObjectsCompareOverviewInfoType
	TRIGGER            ObjectsCompareOverviewInfoType
	SYNONYM            ObjectsCompareOverviewInfoType
	FUNCTION           ObjectsCompareOverviewInfoType
	PROCEDURE          ObjectsCompareOverviewInfoType
	TYPE               ObjectsCompareOverviewInfoType
	RULE               ObjectsCompareOverviewInfoType
	DEFAULT_TYPE       ObjectsCompareOverviewInfoType
	PLAN_GUIDE         ObjectsCompareOverviewInfoType
	CONSTRAINT         ObjectsCompareOverviewInfoType
	FILE_GROUP         ObjectsCompareOverviewInfoType
	PARTITION_FUNCTION ObjectsCompareOverviewInfoType
	PARTITION_SCHEME   ObjectsCompareOverviewInfoType
	TABLE_COLLATION    ObjectsCompareOverviewInfoType
	EXTENSIONS         ObjectsCompareOverviewInfoType
}

func GetObjectsCompareOverviewInfoTypeEnum() ObjectsCompareOverviewInfoTypeEnum {
	return ObjectsCompareOverviewInfoTypeEnum{
		DB: ObjectsCompareOverviewInfoType{
			value: "DB",
		},
		TABLE: ObjectsCompareOverviewInfoType{
			value: "TABLE",
		},
		VIEW: ObjectsCompareOverviewInfoType{
			value: "VIEW",
		},
		EVENT: ObjectsCompareOverviewInfoType{
			value: "EVENT",
		},
		ROUTINE: ObjectsCompareOverviewInfoType{
			value: "ROUTINE",
		},
		INDEX: ObjectsCompareOverviewInfoType{
			value: "INDEX",
		},
		TRIGGER: ObjectsCompareOverviewInfoType{
			value: "TRIGGER",
		},
		SYNONYM: ObjectsCompareOverviewInfoType{
			value: "SYNONYM",
		},
		FUNCTION: ObjectsCompareOverviewInfoType{
			value: "FUNCTION",
		},
		PROCEDURE: ObjectsCompareOverviewInfoType{
			value: "PROCEDURE",
		},
		TYPE: ObjectsCompareOverviewInfoType{
			value: "TYPE",
		},
		RULE: ObjectsCompareOverviewInfoType{
			value: "RULE",
		},
		DEFAULT_TYPE: ObjectsCompareOverviewInfoType{
			value: "DEFAULT_TYPE",
		},
		PLAN_GUIDE: ObjectsCompareOverviewInfoType{
			value: "PLAN_GUIDE",
		},
		CONSTRAINT: ObjectsCompareOverviewInfoType{
			value: "CONSTRAINT",
		},
		FILE_GROUP: ObjectsCompareOverviewInfoType{
			value: "FILE_GROUP",
		},
		PARTITION_FUNCTION: ObjectsCompareOverviewInfoType{
			value: "PARTITION_FUNCTION",
		},
		PARTITION_SCHEME: ObjectsCompareOverviewInfoType{
			value: "PARTITION_SCHEME",
		},
		TABLE_COLLATION: ObjectsCompareOverviewInfoType{
			value: "TABLE_COLLATION",
		},
		EXTENSIONS: ObjectsCompareOverviewInfoType{
			value: "EXTENSIONS",
		},
	}
}

func (c ObjectsCompareOverviewInfoType) Value() string {
	return c.value
}

func (c ObjectsCompareOverviewInfoType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ObjectsCompareOverviewInfoType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}

type ObjectsCompareOverviewInfoStatus struct {
	value string
}

type ObjectsCompareOverviewInfoStatusEnum struct {
	CONSISTENT             ObjectsCompareOverviewInfoStatus
	INCONSISTENT           ObjectsCompareOverviewInfoStatus
	COMPARING              ObjectsCompareOverviewInfoStatus
	WAITING_FOR_COMPARISON ObjectsCompareOverviewInfoStatus
	FAILED_TO_COMPARE      ObjectsCompareOverviewInfoStatus
	TARGET_DB_NOT_EXIST    ObjectsCompareOverviewInfoStatus
	CAN_NOT_COMPARE        ObjectsCompareOverviewInfoStatus
}

func GetObjectsCompareOverviewInfoStatusEnum() ObjectsCompareOverviewInfoStatusEnum {
	return ObjectsCompareOverviewInfoStatusEnum{
		CONSISTENT: ObjectsCompareOverviewInfoStatus{
			value: "CONSISTENT",
		},
		INCONSISTENT: ObjectsCompareOverviewInfoStatus{
			value: "INCONSISTENT",
		},
		COMPARING: ObjectsCompareOverviewInfoStatus{
			value: "COMPARING",
		},
		WAITING_FOR_COMPARISON: ObjectsCompareOverviewInfoStatus{
			value: "WAITING_FOR_COMPARISON",
		},
		FAILED_TO_COMPARE: ObjectsCompareOverviewInfoStatus{
			value: "FAILED_TO_COMPARE",
		},
		TARGET_DB_NOT_EXIST: ObjectsCompareOverviewInfoStatus{
			value: "TARGET_DB_NOT_EXIST",
		},
		CAN_NOT_COMPARE: ObjectsCompareOverviewInfoStatus{
			value: "CAN_NOT_COMPARE",
		},
	}
}

func (c ObjectsCompareOverviewInfoStatus) Value() string {
	return c.value
}

func (c ObjectsCompareOverviewInfoStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ObjectsCompareOverviewInfoStatus) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
