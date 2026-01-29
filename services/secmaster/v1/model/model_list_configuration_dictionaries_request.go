package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListConfigurationDictionariesRequest Request Object
type ListConfigurationDictionariesRequest struct {

	// 用户当前语言环境
	XLanguage string `json:"X-Language"`

	// 字典使用场景，告警为ALERT，等级云为QUAD_CLOUD
	Scene *string `json:"scene,omitempty"`

	// 仅QUAD_CLOUD场景使用[S1, S2, S3, S4]
	Level *string `json:"level,omitempty"`

	// 该字典作用域
	Scope *string `json:"scope,omitempty"`

	// 字典key
	DictKey *string `json:"dict_key,omitempty"`

	// 是否为系统内置字典数据,true：系统内置，false：自定义
	IsBuiltIn *bool `json:"is_built_in,omitempty"`

	// 分页参数：返回起始偏移量
	Offset *int32 `json:"offset,omitempty"`

	// 分页参数：返回数据量大小
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListConfigurationDictionariesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListConfigurationDictionariesRequest struct{}"
	}

	return strings.Join([]string{"ListConfigurationDictionariesRequest", string(data)}, " ")
}
