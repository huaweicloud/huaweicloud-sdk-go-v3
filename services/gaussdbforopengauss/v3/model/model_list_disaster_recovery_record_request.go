package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListDisasterRecoveryRecordRequest Request Object
type ListDisasterRecoveryRecordRequest struct {

	// 语言。
	XLanguage *ListDisasterRecoveryRecordRequestXLanguage `json:"X-Language,omitempty"`

	// 实例id。
	InstanceId string `json:"instance_id"`

	// 实体id（容灾id）
	EntityId string `json:"entity_id"`

	// 实体类型（容灾类型）
	EntityType ListDisasterRecoveryRecordRequestEntityType `json:"entity_type"`

	// 查询记录数。默认为100，不能为负数，最小值为1，最大值为100。
	Limit *int32 `json:"limit,omitempty"`

	// 索引位置，偏移量。从第一条数据偏移offset条数据后开始查询，默认为0（偏移0条数据，表示从第一条数据开始查询），必须为数字，不能为负数。
	Offset *int32 `json:"offset,omitempty"`
}

func (o ListDisasterRecoveryRecordRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDisasterRecoveryRecordRequest struct{}"
	}

	return strings.Join([]string{"ListDisasterRecoveryRecordRequest", string(data)}, " ")
}

type ListDisasterRecoveryRecordRequestXLanguage struct {
	value string
}

type ListDisasterRecoveryRecordRequestXLanguageEnum struct {
	ZH_CN ListDisasterRecoveryRecordRequestXLanguage
	EN_US ListDisasterRecoveryRecordRequestXLanguage
}

func GetListDisasterRecoveryRecordRequestXLanguageEnum() ListDisasterRecoveryRecordRequestXLanguageEnum {
	return ListDisasterRecoveryRecordRequestXLanguageEnum{
		ZH_CN: ListDisasterRecoveryRecordRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ListDisasterRecoveryRecordRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ListDisasterRecoveryRecordRequestXLanguage) Value() string {
	return c.value
}

func (c ListDisasterRecoveryRecordRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListDisasterRecoveryRecordRequestXLanguage) UnmarshalJSON(b []byte) error {
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

type ListDisasterRecoveryRecordRequestEntityType struct {
	value string
}

type ListDisasterRecoveryRecordRequestEntityTypeEnum struct {
	DR ListDisasterRecoveryRecordRequestEntityType
}

func GetListDisasterRecoveryRecordRequestEntityTypeEnum() ListDisasterRecoveryRecordRequestEntityTypeEnum {
	return ListDisasterRecoveryRecordRequestEntityTypeEnum{
		DR: ListDisasterRecoveryRecordRequestEntityType{
			value: "dr",
		},
	}
}

func (c ListDisasterRecoveryRecordRequestEntityType) Value() string {
	return c.value
}

func (c ListDisasterRecoveryRecordRequestEntityType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListDisasterRecoveryRecordRequestEntityType) UnmarshalJSON(b []byte) error {
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
