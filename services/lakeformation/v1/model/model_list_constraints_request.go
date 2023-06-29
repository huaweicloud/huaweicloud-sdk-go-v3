package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListConstraintsRequest Request Object
type ListConstraintsRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	// catalog名字
	CatalogName string `json:"catalog_name"`

	// 数据库名字
	DatabaseName string `json:"database_name"`

	// 表名称
	TableName string `json:"table_name"`

	// 列限制条件类型
	Type ListConstraintsRequestType `json:"type"`

	// 获取表外键时使用，指定外键被引用表所在的数据库名
	ParentDb *string `json:"parent_db,omitempty"`

	// 获取表外键时使用，指定外键被引用表的表名
	ParentTbl *string `json:"parent_tbl,omitempty"`
}

func (o ListConstraintsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListConstraintsRequest struct{}"
	}

	return strings.Join([]string{"ListConstraintsRequest", string(data)}, " ")
}

type ListConstraintsRequestType struct {
	value string
}

type ListConstraintsRequestTypeEnum struct {
	CHECK_CSTR    ListConstraintsRequestType
	DEFAULT_CSTR  ListConstraintsRequestType
	NOT_NULL_CSTR ListConstraintsRequestType
	UNIQUE_CSTR   ListConstraintsRequestType
	PRIMARY_KEY   ListConstraintsRequestType
	FOREIGN_KEY   ListConstraintsRequestType
}

func GetListConstraintsRequestTypeEnum() ListConstraintsRequestTypeEnum {
	return ListConstraintsRequestTypeEnum{
		CHECK_CSTR: ListConstraintsRequestType{
			value: "CHECK_CSTR",
		},
		DEFAULT_CSTR: ListConstraintsRequestType{
			value: "DEFAULT_CSTR",
		},
		NOT_NULL_CSTR: ListConstraintsRequestType{
			value: "NOT_NULL_CSTR",
		},
		UNIQUE_CSTR: ListConstraintsRequestType{
			value: "UNIQUE_CSTR",
		},
		PRIMARY_KEY: ListConstraintsRequestType{
			value: "PRIMARY_KEY",
		},
		FOREIGN_KEY: ListConstraintsRequestType{
			value: "FOREIGN_KEY",
		},
	}
}

func (c ListConstraintsRequestType) Value() string {
	return c.value
}

func (c ListConstraintsRequestType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListConstraintsRequestType) UnmarshalJSON(b []byte) error {
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
