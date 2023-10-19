package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowConfigTemplateResponse Response Object
type ShowConfigTemplateResponse struct {

	// 模板ID
	TemplateId *string `json:"template_id,omitempty"`

	// 缓存实例类型。取值范围如下： - single：表示单机实例 - ha：表示主备实例 - cluster：表示cluster集群实例 - proxy：表示Proxy集群实例 [- ha_rw_split： 表示读写分离实例](tag:hws)
	CacheMode *string `json:"cache_mode,omitempty"`

	// 模板的描述信息
	Description *string `json:"description,omitempty"`

	// 缓存引擎：Redis[和Memcached](tag:hws,hws_hk,ocb,sbc,tm,ctc,cmcc)。
	Engine *string `json:"engine,omitempty"`

	// 缓存版本。  当缓存引擎为Redis时，取值为4.0或5.0。  [当缓存引擎为Memcached时，该字段为可选，取值为空。](tag:hws,hws_hk,ocb,sbc,tm,ctc,cmcc)
	EngineVersion *string `json:"engine_version,omitempty"`

	// 模板名称
	Name *string `json:"name,omitempty"`

	// 产品类型。 取值范围如下： - generic：普通版 - enterprise：企业版
	ProductType *ShowConfigTemplateResponseProductType `json:"product_type,omitempty"`

	// 存储类型。 取值范围如下： - DRAM - SSD
	StorageType *ShowConfigTemplateResponseStorageType `json:"storage_type,omitempty"`

	// 模板类型
	Type *string `json:"type,omitempty"`

	// 模板创建时间，仅在自定义参数模板中有意义，格式例如：2023-05-10T11:09:35.802Z
	CreatedAt      *string `json:"created_at,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowConfigTemplateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowConfigTemplateResponse struct{}"
	}

	return strings.Join([]string{"ShowConfigTemplateResponse", string(data)}, " ")
}

type ShowConfigTemplateResponseProductType struct {
	value string
}

type ShowConfigTemplateResponseProductTypeEnum struct {
	GENERIC    ShowConfigTemplateResponseProductType
	ENTERPRISE ShowConfigTemplateResponseProductType
}

func GetShowConfigTemplateResponseProductTypeEnum() ShowConfigTemplateResponseProductTypeEnum {
	return ShowConfigTemplateResponseProductTypeEnum{
		GENERIC: ShowConfigTemplateResponseProductType{
			value: "generic",
		},
		ENTERPRISE: ShowConfigTemplateResponseProductType{
			value: "enterprise",
		},
	}
}

func (c ShowConfigTemplateResponseProductType) Value() string {
	return c.value
}

func (c ShowConfigTemplateResponseProductType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowConfigTemplateResponseProductType) UnmarshalJSON(b []byte) error {
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

type ShowConfigTemplateResponseStorageType struct {
	value string
}

type ShowConfigTemplateResponseStorageTypeEnum struct {
	DRAM ShowConfigTemplateResponseStorageType
	SSD  ShowConfigTemplateResponseStorageType
}

func GetShowConfigTemplateResponseStorageTypeEnum() ShowConfigTemplateResponseStorageTypeEnum {
	return ShowConfigTemplateResponseStorageTypeEnum{
		DRAM: ShowConfigTemplateResponseStorageType{
			value: "DRAM",
		},
		SSD: ShowConfigTemplateResponseStorageType{
			value: "SSD",
		},
	}
}

func (c ShowConfigTemplateResponseStorageType) Value() string {
	return c.value
}

func (c ShowConfigTemplateResponseStorageType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowConfigTemplateResponseStorageType) UnmarshalJSON(b []byte) error {
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
