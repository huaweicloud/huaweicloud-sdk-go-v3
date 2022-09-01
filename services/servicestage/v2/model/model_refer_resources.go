package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 部署资源。
type ReferResources struct {

	// 资源ID。
	Id *string `json:"id,omitempty" xml:"id"`

	Type *ResourceType `json:"type,omitempty" xml:"type"`

	// 应用别名，dcs时才提供，支持“distributed_session”、“distributed_cache”、“distributed_session, distributed_cache”，  默认值是“distributed_session, distributed_cache”。
	ReferAlias *string `json:"refer_alias,omitempty" xml:"refer_alias"`

	// 引用资源参数。
	Parameters *interface{} `json:"parameters,omitempty" xml:"parameters"`
}

func (o ReferResources) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ReferResources struct{}"
	}

	return strings.Join([]string{"ReferResources", string(data)}, " ")
}
