package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type CatalogTargetInfo struct {

	// 事件目标分类ID
	Id *string `json:"id,omitempty"`

	// 事件目标分类名称
	Name *string `json:"name,omitempty"`

	// 事件目标分类名称展示说明
	Label *string `json:"label,omitempty"`

	// 事件目标分类描述
	Description *string `json:"description,omitempty"`

	// 提供方类型，OFFICIAL：官方云服务事件目标；CUSTOM：第三方事件目标
	ProviderType *CatalogTargetInfoProviderType `json:"provider_type,omitempty"`

	Parameters *[]interface{} `json:"parameters,omitempty"`

	// 创建UTC时间
	CreatedTime *string `json:"created_time,omitempty"`

	// 更新UTC时间
	UpdatedTime *string `json:"updated_time,omitempty"`
}

func (o CatalogTargetInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CatalogTargetInfo struct{}"
	}

	return strings.Join([]string{"CatalogTargetInfo", string(data)}, " ")
}

type CatalogTargetInfoProviderType struct {
	value string
}

type CatalogTargetInfoProviderTypeEnum struct {
	OFFICIAL CatalogTargetInfoProviderType
	CUSTOM   CatalogTargetInfoProviderType
}

func GetCatalogTargetInfoProviderTypeEnum() CatalogTargetInfoProviderTypeEnum {
	return CatalogTargetInfoProviderTypeEnum{
		OFFICIAL: CatalogTargetInfoProviderType{
			value: "OFFICIAL",
		},
		CUSTOM: CatalogTargetInfoProviderType{
			value: "CUSTOM",
		},
	}
}

func (c CatalogTargetInfoProviderType) Value() string {
	return c.value
}

func (c CatalogTargetInfoProviderType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CatalogTargetInfoProviderType) UnmarshalJSON(b []byte) error {
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
