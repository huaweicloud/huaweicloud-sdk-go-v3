package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListSpaceAnalysisRequest Request Object
type ListSpaceAnalysisRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	// 语言
	XLanguage *string `json:"X-Language,omitempty"`

	// 对象类型
	ObjectType ListSpaceAnalysisRequestObjectType `json:"object_type"`

	// 数据库ID
	DatabaseId *string `json:"database_id,omitempty"`

	// 偏移量。从第一条数据偏移offset条数据后开始查询，默认为0（偏移0条数据，表示从第一条数据开始查询），必须为数字，不能为负数。offset必须是limit的整数倍。
	Offset *string `json:"offset,omitempty"`

	// 查询记录数。默认为100，不能为负数，最小值为1，最大值为100。
	Limit *string `json:"limit,omitempty"`

	// 是否返回实例级别数据，取值：true或者false
	ShowInstanceInfo *string `json:"show_instance_info,omitempty"`

	// 引擎类型
	DatastoreType ListSpaceAnalysisRequestDatastoreType `json:"datastore_type"`
}

func (o ListSpaceAnalysisRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSpaceAnalysisRequest struct{}"
	}

	return strings.Join([]string{"ListSpaceAnalysisRequest", string(data)}, " ")
}

type ListSpaceAnalysisRequestObjectType struct {
	value string
}

type ListSpaceAnalysisRequestObjectTypeEnum struct {
	DATABASE ListSpaceAnalysisRequestObjectType
	TABLE    ListSpaceAnalysisRequestObjectType
}

func GetListSpaceAnalysisRequestObjectTypeEnum() ListSpaceAnalysisRequestObjectTypeEnum {
	return ListSpaceAnalysisRequestObjectTypeEnum{
		DATABASE: ListSpaceAnalysisRequestObjectType{
			value: "database",
		},
		TABLE: ListSpaceAnalysisRequestObjectType{
			value: "table",
		},
	}
}

func (c ListSpaceAnalysisRequestObjectType) Value() string {
	return c.value
}

func (c ListSpaceAnalysisRequestObjectType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListSpaceAnalysisRequestObjectType) UnmarshalJSON(b []byte) error {
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

type ListSpaceAnalysisRequestDatastoreType struct {
	value string
}

type ListSpaceAnalysisRequestDatastoreTypeEnum struct {
	MY_SQL              ListSpaceAnalysisRequestDatastoreType
	GAUSS_DB_FOR_MY_SQL ListSpaceAnalysisRequestDatastoreType
	SQL_SERVER          ListSpaceAnalysisRequestDatastoreType
}

func GetListSpaceAnalysisRequestDatastoreTypeEnum() ListSpaceAnalysisRequestDatastoreTypeEnum {
	return ListSpaceAnalysisRequestDatastoreTypeEnum{
		MY_SQL: ListSpaceAnalysisRequestDatastoreType{
			value: "MySQL",
		},
		GAUSS_DB_FOR_MY_SQL: ListSpaceAnalysisRequestDatastoreType{
			value: "GaussDB(for MySQL)",
		},
		SQL_SERVER: ListSpaceAnalysisRequestDatastoreType{
			value: "SQLServer",
		},
	}
}

func (c ListSpaceAnalysisRequestDatastoreType) Value() string {
	return c.value
}

func (c ListSpaceAnalysisRequestDatastoreType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListSpaceAnalysisRequestDatastoreType) UnmarshalJSON(b []byte) error {
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
