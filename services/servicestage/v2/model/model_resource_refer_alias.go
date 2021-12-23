package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 应用别名，dcs时才提供，支持“distributed_session”、“distributed_cache”、“distributed_session, distributed_cache”，  默认值是“distributed_session, distributed_cache”。
type ResourceReferAlias struct {
}

func (o ResourceReferAlias) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceReferAlias struct{}"
	}

	return strings.Join([]string{"ResourceReferAlias", string(data)}, " ")
}
