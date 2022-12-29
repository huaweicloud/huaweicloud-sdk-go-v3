package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 配额信息，包括总配额和已使用配额。
type QuotaInfo struct {

	// 资源类型。  取值： loadbalancer、listener、ipgroup、pool、member、members_per_pool、 healthmonitor、l7policy、certificate、security_policy、 ipgroup_bindings、ipgroup_max_length。  members_per_pool表示一个pool下最多可关联的member数量。  ipgroup_bindings表示一个ipgroup下最多可关联的listener数量。  ipgroup_max_length表示一个ipgroup下最多设置的ip地址数量。
	QuotaKey string `json:"quota_key"`

	// 总配额。  取值： - 大于等于0：表示当前配额数量。 - -1：表示无配额限制。
	QuotaLimit int32 `json:"quota_limit"`

	// 已使用配额。
	Used int32 `json:"used"`

	// 配额单位。  取值：count，表示个数。
	Unit string `json:"unit"`
}

func (o QuotaInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QuotaInfo struct{}"
	}

	return strings.Join([]string{"QuotaInfo", string(data)}, " ")
}
