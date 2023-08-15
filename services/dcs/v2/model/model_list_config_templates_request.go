package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListConfigTemplatesRequest Request Object
type ListConfigTemplatesRequest struct {

	// 参数模板名称，支持模糊查找
	Name *string `json:"name,omitempty"`

	// 模板ID
	TemplateId *string `json:"template_id,omitempty"`

	// 模板类型
	Type ListConfigTemplatesRequestType `json:"type"`

	// 缓存引擎：Redis[和Memcached](tag:hws,hws_hk,ocb,sbc,tm,ctc,cmcc)。
	Engine *string `json:"engine,omitempty"`

	// 缓存版本。  当缓存引擎为Redis时，取值为[3.0/4.0/5.0](tag:ctc,cmc)[3.0/4.0/5.0/6.0](tag:ocb,otc,sbc,g42,tm)[4.0/5.0/6.0](tag:hws,hws_hk)。  [当缓存引擎为Memcached时，该字段为可选，取值为空。](tag:hws,hws_hk,ocb,sbc,tm,ctc,cmcc)
	EngineVersion *string `json:"engine_version,omitempty"`

	// 缓存实例类型。取值范围如下： - single：表示单机实例 - ha：表示主备实例 - cluster：表示cluster集群实例 - proxy：表示Proxy集群实例 [- ha_rw_split：表示读写分离实例](tag:hws)
	CacheMode *string `json:"cache_mode,omitempty"`

	// 模板的描述信息
	Description *string `json:"description,omitempty"`

	// 偏移量，表示从此偏移量开始查询， offset大于等于0
	Offset *int32 `json:"offset,omitempty"`

	// 每页显示条数，最小值为1，最大值为1000，若不设置该参数，则为10。
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListConfigTemplatesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListConfigTemplatesRequest struct{}"
	}

	return strings.Join([]string{"ListConfigTemplatesRequest", string(data)}, " ")
}

type ListConfigTemplatesRequestType struct {
	value string
}

type ListConfigTemplatesRequestTypeEnum struct {
	SYS  ListConfigTemplatesRequestType
	USER ListConfigTemplatesRequestType
}

func GetListConfigTemplatesRequestTypeEnum() ListConfigTemplatesRequestTypeEnum {
	return ListConfigTemplatesRequestTypeEnum{
		SYS: ListConfigTemplatesRequestType{
			value: "sys",
		},
		USER: ListConfigTemplatesRequestType{
			value: "user",
		},
	}
}

func (c ListConfigTemplatesRequestType) Value() string {
	return c.value
}

func (c ListConfigTemplatesRequestType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListConfigTemplatesRequestType) UnmarshalJSON(b []byte) error {
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
