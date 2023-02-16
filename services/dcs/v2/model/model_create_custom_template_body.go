package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateCustomTemplateBody struct {

	// 来源系统模板ID
	TemplateId *string `json:"template_id,omitempty"`

	// 模板名称
	Name string `json:"name"`

	// 模板类型
	Type *string `json:"type,omitempty"`

	// 缓存引擎：Redis[和Memcached](tag:hws,hws_hk,ocb,sbc,tm,ctc,cmcc)。
	Engine *string `json:"engine,omitempty"`

	// 缓存实例类型。取值范围如下： - single：表示单机实例 - ha：表示主备实例 - cluster：表示cluster集群实例 - proxy：表示Proxy集群实例 [- ha_rw_split： 表示读写分离实例](tag:hws)
	CacheMode *string `json:"cache_mode,omitempty"`

	// 模板的描述信息
	Description *string `json:"description,omitempty"`

	// 缓存版本。  当缓存引擎为Redis时，取值为4.0或5.0。  [当缓存引擎为Memcached时，该字段为可选，取值为空。](tag:hws,hws_hk,ocb,sbc,tm,ctc,cmcc)
	EngineVersion *string `json:"engine_version,omitempty"`

	// 参数配置信息
	Params map[string]string `json:"params,omitempty"`
}

func (o CreateCustomTemplateBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCustomTemplateBody struct{}"
	}

	return strings.Join([]string{"CreateCustomTemplateBody", string(data)}, " ")
}
