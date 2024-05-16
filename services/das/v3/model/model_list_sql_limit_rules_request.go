package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListSqlLimitRulesRequest Request Object
type ListSqlLimitRulesRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	// 偏移量。从第一条数据偏移offset条数据后开始查询，默认为0（偏移0条数据，表示从第一条数据开始查询），必须为数字，不能为负数。
	Offset *int32 `json:"offset,omitempty"`

	// 查询记录数。默认为100，不能为负数，最小值为1，最大值为100。
	Limit *int32 `json:"limit,omitempty"`

	// 数据库类型
	DatastoreType ListSqlLimitRulesRequestDatastoreType `json:"datastore_type"`

	// 数据库名（PostgreSQL必填）
	DatabaseName *string `json:"database_name,omitempty"`

	// 语言
	XLanguage *ListSqlLimitRulesRequestXLanguage `json:"X-Language,omitempty"`
}

func (o ListSqlLimitRulesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSqlLimitRulesRequest struct{}"
	}

	return strings.Join([]string{"ListSqlLimitRulesRequest", string(data)}, " ")
}

type ListSqlLimitRulesRequestDatastoreType struct {
	value string
}

type ListSqlLimitRulesRequestDatastoreTypeEnum struct {
	MY_SQL      ListSqlLimitRulesRequestDatastoreType
	POSTGRE_SQL ListSqlLimitRulesRequestDatastoreType
}

func GetListSqlLimitRulesRequestDatastoreTypeEnum() ListSqlLimitRulesRequestDatastoreTypeEnum {
	return ListSqlLimitRulesRequestDatastoreTypeEnum{
		MY_SQL: ListSqlLimitRulesRequestDatastoreType{
			value: "MySQL",
		},
		POSTGRE_SQL: ListSqlLimitRulesRequestDatastoreType{
			value: "PostgreSQL",
		},
	}
}

func (c ListSqlLimitRulesRequestDatastoreType) Value() string {
	return c.value
}

func (c ListSqlLimitRulesRequestDatastoreType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListSqlLimitRulesRequestDatastoreType) UnmarshalJSON(b []byte) error {
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

type ListSqlLimitRulesRequestXLanguage struct {
	value string
}

type ListSqlLimitRulesRequestXLanguageEnum struct {
	ZH_CN ListSqlLimitRulesRequestXLanguage
	EN_US ListSqlLimitRulesRequestXLanguage
}

func GetListSqlLimitRulesRequestXLanguageEnum() ListSqlLimitRulesRequestXLanguageEnum {
	return ListSqlLimitRulesRequestXLanguageEnum{
		ZH_CN: ListSqlLimitRulesRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ListSqlLimitRulesRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ListSqlLimitRulesRequestXLanguage) Value() string {
	return c.value
}

func (c ListSqlLimitRulesRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListSqlLimitRulesRequestXLanguage) UnmarshalJSON(b []byte) error {
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
