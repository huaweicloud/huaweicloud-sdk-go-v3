package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 部署资源。
type ReferResourceCreate struct {
	// 资源ID。

	Id string `json:"id"`

	Type *ResourceType `json:"type"`
	// 应用别名，dcs时才提供，支持“distributed_session”、“distributed_cache”、“distributed_session, distributed_cache”，  默认值是“distributed_session, distributed_cache”。

	ReferAlias *string `json:"refer_alias,omitempty"`
	// 引用资源参数。

	Parameters *interface{} `json:"parameters,omitempty"`
}

func (o ReferResourceCreate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ReferResourceCreate struct{}"
	}

	return strings.Join([]string{"ReferResourceCreate", string(data)}, " ")
}
