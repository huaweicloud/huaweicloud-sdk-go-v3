package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListTransactionsRequest Request Object
type ListTransactionsRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	// 数据库类型。仅支持MySQL
	DatastoreType ListTransactionsRequestDatastoreType `json:"datastore_type"`

	// 语言
	XLanguage *ListTransactionsRequestXLanguage `json:"X-Language,omitempty"`

	// 采集开始时间（Unix timestamp），单位：毫秒。
	StartAt int64 `json:"start_at"`

	// 采集结束时间（Unix timestamp），单位：毫秒。
	EndAt int64 `json:"end_at"`

	// 页数
	PageNum *int32 `json:"page_num,omitempty"`

	// 页大小
	PageSize *int32 `json:"page_size,omitempty"`

	// 排序字段
	Order *ListTransactionsRequestOrder `json:"order,omitempty"`

	// 升序|降序
	OrderBy *ListTransactionsRequestOrderBy `json:"order_by,omitempty"`

	// 持续时间下限
	LastSecMin *int64 `json:"last_sec_min,omitempty"`

	// 持续时间上限
	LastSecMax *int64 `json:"last_sec_max,omitempty"`
}

func (o ListTransactionsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTransactionsRequest struct{}"
	}

	return strings.Join([]string{"ListTransactionsRequest", string(data)}, " ")
}

type ListTransactionsRequestDatastoreType struct {
	value string
}

type ListTransactionsRequestDatastoreTypeEnum struct {
	MY_SQL ListTransactionsRequestDatastoreType
}

func GetListTransactionsRequestDatastoreTypeEnum() ListTransactionsRequestDatastoreTypeEnum {
	return ListTransactionsRequestDatastoreTypeEnum{
		MY_SQL: ListTransactionsRequestDatastoreType{
			value: "MySQL",
		},
	}
}

func (c ListTransactionsRequestDatastoreType) Value() string {
	return c.value
}

func (c ListTransactionsRequestDatastoreType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListTransactionsRequestDatastoreType) UnmarshalJSON(b []byte) error {
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

type ListTransactionsRequestXLanguage struct {
	value string
}

type ListTransactionsRequestXLanguageEnum struct {
	ZH_CN ListTransactionsRequestXLanguage
	EN_US ListTransactionsRequestXLanguage
}

func GetListTransactionsRequestXLanguageEnum() ListTransactionsRequestXLanguageEnum {
	return ListTransactionsRequestXLanguageEnum{
		ZH_CN: ListTransactionsRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ListTransactionsRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ListTransactionsRequestXLanguage) Value() string {
	return c.value
}

func (c ListTransactionsRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListTransactionsRequestXLanguage) UnmarshalJSON(b []byte) error {
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

type ListTransactionsRequestOrder struct {
	value string
}

type ListTransactionsRequestOrderEnum struct {
	OCCURRENCE_TIME        ListTransactionsRequestOrder
	LAST_SEC               ListTransactionsRequestOrder
	WAIT_LOCK_STRUCT_COUNT ListTransactionsRequestOrder
	HOLD_LOCK_STRUCT_COUNT ListTransactionsRequestOrder
	COLLECT_TIME           ListTransactionsRequestOrder
}

func GetListTransactionsRequestOrderEnum() ListTransactionsRequestOrderEnum {
	return ListTransactionsRequestOrderEnum{
		OCCURRENCE_TIME: ListTransactionsRequestOrder{
			value: "occurrenceTime",
		},
		LAST_SEC: ListTransactionsRequestOrder{
			value: "lastSec",
		},
		WAIT_LOCK_STRUCT_COUNT: ListTransactionsRequestOrder{
			value: "waitLockStructCount",
		},
		HOLD_LOCK_STRUCT_COUNT: ListTransactionsRequestOrder{
			value: "holdLockStructCount",
		},
		COLLECT_TIME: ListTransactionsRequestOrder{
			value: "collectTime",
		},
	}
}

func (c ListTransactionsRequestOrder) Value() string {
	return c.value
}

func (c ListTransactionsRequestOrder) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListTransactionsRequestOrder) UnmarshalJSON(b []byte) error {
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

type ListTransactionsRequestOrderBy struct {
	value string
}

type ListTransactionsRequestOrderByEnum struct {
	ASC  ListTransactionsRequestOrderBy
	DESC ListTransactionsRequestOrderBy
}

func GetListTransactionsRequestOrderByEnum() ListTransactionsRequestOrderByEnum {
	return ListTransactionsRequestOrderByEnum{
		ASC: ListTransactionsRequestOrderBy{
			value: "asc",
		},
		DESC: ListTransactionsRequestOrderBy{
			value: "desc",
		},
	}
}

func (c ListTransactionsRequestOrderBy) Value() string {
	return c.value
}

func (c ListTransactionsRequestOrderBy) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListTransactionsRequestOrderBy) UnmarshalJSON(b []byte) error {
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
