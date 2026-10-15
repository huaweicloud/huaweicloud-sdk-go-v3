package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListHotKeysRequest Request Object
type ListHotKeysRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	// 统计维度。qps-按QPS，traffic-按流量
	StatType ListHotKeysRequestStatType `json:"stat_type"`

	// 开始时间（毫秒时间戳）
	StartAt int64 `json:"start_at"`

	// 结束时间（毫秒时间戳）
	EndAt int64 `json:"end_at"`

	// key类型过滤（多选）。取值：string、hash、list、zset、set、stream、exhash
	KeyTypes *[]ListHotKeysRequestKeyTypes `json:"key_types,omitempty"`

	// key名称过滤（模糊匹配）
	KeyName *string `json:"key_name,omitempty"`

	// 数据库ID过滤
	DbId *string `json:"db_id,omitempty"`

	// 偏移量，默认0
	Offset *int32 `json:"offset,omitempty"`

	// 每页条数，默认20，最大1000
	Limit *int32 `json:"limit,omitempty"`

	// 排序字段。qps/traffic_in/traffic_out，默认qps
	SortField *ListHotKeysRequestSortField `json:"sort_field,omitempty"`

	// 是否升序，默认false（降序）
	Asc *bool `json:"asc,omitempty"`

	// 请求语言类型。en-us：英文。zh-cn：中文。
	XLanguage *ListHotKeysRequestXLanguage `json:"X-Language,omitempty"`
}

func (o ListHotKeysRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListHotKeysRequest struct{}"
	}

	return strings.Join([]string{"ListHotKeysRequest", string(data)}, " ")
}

type ListHotKeysRequestStatType struct {
	value string
}

type ListHotKeysRequestStatTypeEnum struct {
	QPS     ListHotKeysRequestStatType
	TRAFFIC ListHotKeysRequestStatType
}

func GetListHotKeysRequestStatTypeEnum() ListHotKeysRequestStatTypeEnum {
	return ListHotKeysRequestStatTypeEnum{
		QPS: ListHotKeysRequestStatType{
			value: "qps",
		},
		TRAFFIC: ListHotKeysRequestStatType{
			value: "traffic",
		},
	}
}

func (c ListHotKeysRequestStatType) Value() string {
	return c.value
}

func (c ListHotKeysRequestStatType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListHotKeysRequestStatType) UnmarshalJSON(b []byte) error {
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

type ListHotKeysRequestKeyTypes struct {
	value string
}

type ListHotKeysRequestKeyTypesEnum struct {
	STRING ListHotKeysRequestKeyTypes
	HASH   ListHotKeysRequestKeyTypes
	LIST   ListHotKeysRequestKeyTypes
	ZSET   ListHotKeysRequestKeyTypes
	SET    ListHotKeysRequestKeyTypes
	STREAM ListHotKeysRequestKeyTypes
	EXHASH ListHotKeysRequestKeyTypes
}

func GetListHotKeysRequestKeyTypesEnum() ListHotKeysRequestKeyTypesEnum {
	return ListHotKeysRequestKeyTypesEnum{
		STRING: ListHotKeysRequestKeyTypes{
			value: "string",
		},
		HASH: ListHotKeysRequestKeyTypes{
			value: "hash",
		},
		LIST: ListHotKeysRequestKeyTypes{
			value: "list",
		},
		ZSET: ListHotKeysRequestKeyTypes{
			value: "zset",
		},
		SET: ListHotKeysRequestKeyTypes{
			value: "set",
		},
		STREAM: ListHotKeysRequestKeyTypes{
			value: "stream",
		},
		EXHASH: ListHotKeysRequestKeyTypes{
			value: "exhash",
		},
	}
}

func (c ListHotKeysRequestKeyTypes) Value() string {
	return c.value
}

func (c ListHotKeysRequestKeyTypes) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListHotKeysRequestKeyTypes) UnmarshalJSON(b []byte) error {
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

type ListHotKeysRequestSortField struct {
	value string
}

type ListHotKeysRequestSortFieldEnum struct {
	QPS         ListHotKeysRequestSortField
	TRAFFIC_IN  ListHotKeysRequestSortField
	TRAFFIC_OUT ListHotKeysRequestSortField
}

func GetListHotKeysRequestSortFieldEnum() ListHotKeysRequestSortFieldEnum {
	return ListHotKeysRequestSortFieldEnum{
		QPS: ListHotKeysRequestSortField{
			value: "qps",
		},
		TRAFFIC_IN: ListHotKeysRequestSortField{
			value: "traffic_in",
		},
		TRAFFIC_OUT: ListHotKeysRequestSortField{
			value: "traffic_out",
		},
	}
}

func (c ListHotKeysRequestSortField) Value() string {
	return c.value
}

func (c ListHotKeysRequestSortField) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListHotKeysRequestSortField) UnmarshalJSON(b []byte) error {
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

type ListHotKeysRequestXLanguage struct {
	value string
}

type ListHotKeysRequestXLanguageEnum struct {
	ZH_CN ListHotKeysRequestXLanguage
	EN_US ListHotKeysRequestXLanguage
}

func GetListHotKeysRequestXLanguageEnum() ListHotKeysRequestXLanguageEnum {
	return ListHotKeysRequestXLanguageEnum{
		ZH_CN: ListHotKeysRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ListHotKeysRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ListHotKeysRequestXLanguage) Value() string {
	return c.value
}

func (c ListHotKeysRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListHotKeysRequestXLanguage) UnmarshalJSON(b []byte) error {
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
