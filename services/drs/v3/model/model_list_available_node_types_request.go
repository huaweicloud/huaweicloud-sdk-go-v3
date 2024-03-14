package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListAvailableNodeTypesRequest Request Object
type ListAvailableNodeTypesRequest struct {

	// 请求语言类型。
	XLanguage *ListAvailableNodeTypesRequestXLanguage `json:"X-Language,omitempty"`

	// 引擎类型
	EngineType string `json:"engine_type"`

	// 迁移场景。 - migration:实时迁移 - sync:实时同步 - cloudDataGuard:实时灾备
	DbUseType ListAvailableNodeTypesRequestDbUseType `json:"db_use_type"`

	// 迁移方向，up：入云 ，down：出云，non-dbs：自建。
	JobDirection ListAvailableNodeTypesRequestJobDirection `json:"job_direction"`

	// 是否查询资源售罄情况
	IsUseSelloutInfo *bool `json:"is_use_sellout_info,omitempty"`

	// 是否是双主灾备
	IsMultiWrite *bool `json:"is_multi_write,omitempty"`
}

func (o ListAvailableNodeTypesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAvailableNodeTypesRequest struct{}"
	}

	return strings.Join([]string{"ListAvailableNodeTypesRequest", string(data)}, " ")
}

type ListAvailableNodeTypesRequestXLanguage struct {
	value string
}

type ListAvailableNodeTypesRequestXLanguageEnum struct {
	EN_US ListAvailableNodeTypesRequestXLanguage
	ZH_CN ListAvailableNodeTypesRequestXLanguage
}

func GetListAvailableNodeTypesRequestXLanguageEnum() ListAvailableNodeTypesRequestXLanguageEnum {
	return ListAvailableNodeTypesRequestXLanguageEnum{
		EN_US: ListAvailableNodeTypesRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: ListAvailableNodeTypesRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c ListAvailableNodeTypesRequestXLanguage) Value() string {
	return c.value
}

func (c ListAvailableNodeTypesRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListAvailableNodeTypesRequestXLanguage) UnmarshalJSON(b []byte) error {
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

type ListAvailableNodeTypesRequestDbUseType struct {
	value string
}

type ListAvailableNodeTypesRequestDbUseTypeEnum struct {
	MIGRATION        ListAvailableNodeTypesRequestDbUseType
	SYNC             ListAvailableNodeTypesRequestDbUseType
	CLOUD_DATA_GUARD ListAvailableNodeTypesRequestDbUseType
}

func GetListAvailableNodeTypesRequestDbUseTypeEnum() ListAvailableNodeTypesRequestDbUseTypeEnum {
	return ListAvailableNodeTypesRequestDbUseTypeEnum{
		MIGRATION: ListAvailableNodeTypesRequestDbUseType{
			value: "migration",
		},
		SYNC: ListAvailableNodeTypesRequestDbUseType{
			value: "sync",
		},
		CLOUD_DATA_GUARD: ListAvailableNodeTypesRequestDbUseType{
			value: "cloudDataGuard",
		},
	}
}

func (c ListAvailableNodeTypesRequestDbUseType) Value() string {
	return c.value
}

func (c ListAvailableNodeTypesRequestDbUseType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListAvailableNodeTypesRequestDbUseType) UnmarshalJSON(b []byte) error {
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

type ListAvailableNodeTypesRequestJobDirection struct {
	value string
}

type ListAvailableNodeTypesRequestJobDirectionEnum struct {
	UP      ListAvailableNodeTypesRequestJobDirection
	DOWN    ListAvailableNodeTypesRequestJobDirection
	NON_DBS ListAvailableNodeTypesRequestJobDirection
}

func GetListAvailableNodeTypesRequestJobDirectionEnum() ListAvailableNodeTypesRequestJobDirectionEnum {
	return ListAvailableNodeTypesRequestJobDirectionEnum{
		UP: ListAvailableNodeTypesRequestJobDirection{
			value: "up",
		},
		DOWN: ListAvailableNodeTypesRequestJobDirection{
			value: "down",
		},
		NON_DBS: ListAvailableNodeTypesRequestJobDirection{
			value: "non-dbs",
		},
	}
}

func (c ListAvailableNodeTypesRequestJobDirection) Value() string {
	return c.value
}

func (c ListAvailableNodeTypesRequestJobDirection) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListAvailableNodeTypesRequestJobDirection) UnmarshalJSON(b []byte) error {
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
