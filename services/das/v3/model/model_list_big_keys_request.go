package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListBigKeysRequest Request Object
type ListBigKeysRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	// 统计维度。size-按长度，storage-按存储占用
	StatType ListBigKeysRequestStatType `json:"stat_type"`

	// 开始时间（毫秒时间戳）
	StartAt int64 `json:"start_at"`

	// 结束时间（毫秒时间戳）
	EndAt int64 `json:"end_at"`

	// key类型过滤（多选）。取值：string、hash、list、zset、set、stream、exhash
	KeyTypes *[]ListBigKeysRequestKeyTypes `json:"key_types,omitempty"`

	// key名称过滤（模糊匹配）
	KeyName *string `json:"key_name,omitempty"`

	// 数据库ID过滤
	DbId *string `json:"db_id,omitempty"`

	// 偏移量，默认0
	Offset *int32 `json:"offset,omitempty"`

	// 每页条数，默认20，最大1000
	Limit *int32 `json:"limit,omitempty"`

	// 排序字段，默认size
	SortField *ListBigKeysRequestSortField `json:"sort_field,omitempty"`

	// 是否升序，默认false（降序）
	Asc *bool `json:"asc,omitempty"`

	// 请求语言类型。en-us：英文。zh-cn：中文。
	XLanguage *ListBigKeysRequestXLanguage `json:"X-Language,omitempty"`
}

func (o ListBigKeysRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBigKeysRequest struct{}"
	}

	return strings.Join([]string{"ListBigKeysRequest", string(data)}, " ")
}

type ListBigKeysRequestStatType struct {
	value string
}

type ListBigKeysRequestStatTypeEnum struct {
	SIZE    ListBigKeysRequestStatType
	STORAGE ListBigKeysRequestStatType
}

func GetListBigKeysRequestStatTypeEnum() ListBigKeysRequestStatTypeEnum {
	return ListBigKeysRequestStatTypeEnum{
		SIZE: ListBigKeysRequestStatType{
			value: "size",
		},
		STORAGE: ListBigKeysRequestStatType{
			value: "storage",
		},
	}
}

func (c ListBigKeysRequestStatType) Value() string {
	return c.value
}

func (c ListBigKeysRequestStatType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListBigKeysRequestStatType) UnmarshalJSON(b []byte) error {
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

type ListBigKeysRequestKeyTypes struct {
	value string
}

type ListBigKeysRequestKeyTypesEnum struct {
	STRING ListBigKeysRequestKeyTypes
	HASH   ListBigKeysRequestKeyTypes
	LIST   ListBigKeysRequestKeyTypes
	ZSET   ListBigKeysRequestKeyTypes
	SET    ListBigKeysRequestKeyTypes
	STREAM ListBigKeysRequestKeyTypes
	EXHASH ListBigKeysRequestKeyTypes
}

func GetListBigKeysRequestKeyTypesEnum() ListBigKeysRequestKeyTypesEnum {
	return ListBigKeysRequestKeyTypesEnum{
		STRING: ListBigKeysRequestKeyTypes{
			value: "string",
		},
		HASH: ListBigKeysRequestKeyTypes{
			value: "hash",
		},
		LIST: ListBigKeysRequestKeyTypes{
			value: "list",
		},
		ZSET: ListBigKeysRequestKeyTypes{
			value: "zset",
		},
		SET: ListBigKeysRequestKeyTypes{
			value: "set",
		},
		STREAM: ListBigKeysRequestKeyTypes{
			value: "stream",
		},
		EXHASH: ListBigKeysRequestKeyTypes{
			value: "exhash",
		},
	}
}

func (c ListBigKeysRequestKeyTypes) Value() string {
	return c.value
}

func (c ListBigKeysRequestKeyTypes) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListBigKeysRequestKeyTypes) UnmarshalJSON(b []byte) error {
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

type ListBigKeysRequestSortField struct {
	value string
}

type ListBigKeysRequestSortFieldEnum struct {
	SIZE ListBigKeysRequestSortField
}

func GetListBigKeysRequestSortFieldEnum() ListBigKeysRequestSortFieldEnum {
	return ListBigKeysRequestSortFieldEnum{
		SIZE: ListBigKeysRequestSortField{
			value: "size",
		},
	}
}

func (c ListBigKeysRequestSortField) Value() string {
	return c.value
}

func (c ListBigKeysRequestSortField) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListBigKeysRequestSortField) UnmarshalJSON(b []byte) error {
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

type ListBigKeysRequestXLanguage struct {
	value string
}

type ListBigKeysRequestXLanguageEnum struct {
	ZH_CN ListBigKeysRequestXLanguage
	EN_US ListBigKeysRequestXLanguage
}

func GetListBigKeysRequestXLanguageEnum() ListBigKeysRequestXLanguageEnum {
	return ListBigKeysRequestXLanguageEnum{
		ZH_CN: ListBigKeysRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ListBigKeysRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ListBigKeysRequestXLanguage) Value() string {
	return c.value
}

func (c ListBigKeysRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListBigKeysRequestXLanguage) UnmarshalJSON(b []byte) error {
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
