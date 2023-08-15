package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListLinksRequest Request Object
type ListLinksRequest struct {

	// 请求语言类型。
	XLanguage *ListLinksRequestXLanguage `json:"X-Language,omitempty"`

	// 任务场景。取值： - migration：实时迁移。 - sync：实时同步。 - cloudDataGuard：实时灾备。
	JobType ListLinksRequestJobType `json:"job_type"`

	// 偏移量，表示查询该偏移量后面的记录。
	Offset *int32 `json:"offset,omitempty"`

	// 查询返回记录的数量限制。
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListLinksRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListLinksRequest struct{}"
	}

	return strings.Join([]string{"ListLinksRequest", string(data)}, " ")
}

type ListLinksRequestXLanguage struct {
	value string
}

type ListLinksRequestXLanguageEnum struct {
	EN_US ListLinksRequestXLanguage
	ZH_CN ListLinksRequestXLanguage
}

func GetListLinksRequestXLanguageEnum() ListLinksRequestXLanguageEnum {
	return ListLinksRequestXLanguageEnum{
		EN_US: ListLinksRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: ListLinksRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c ListLinksRequestXLanguage) Value() string {
	return c.value
}

func (c ListLinksRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListLinksRequestXLanguage) UnmarshalJSON(b []byte) error {
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

type ListLinksRequestJobType struct {
	value string
}

type ListLinksRequestJobTypeEnum struct {
	MIGRATION        ListLinksRequestJobType
	SYNC             ListLinksRequestJobType
	CLOUD_DATA_GUARD ListLinksRequestJobType
}

func GetListLinksRequestJobTypeEnum() ListLinksRequestJobTypeEnum {
	return ListLinksRequestJobTypeEnum{
		MIGRATION: ListLinksRequestJobType{
			value: "migration",
		},
		SYNC: ListLinksRequestJobType{
			value: "sync",
		},
		CLOUD_DATA_GUARD: ListLinksRequestJobType{
			value: "cloudDataGuard",
		},
	}
}

func (c ListLinksRequestJobType) Value() string {
	return c.value
}

func (c ListLinksRequestJobType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListLinksRequestJobType) UnmarshalJSON(b []byte) error {
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
