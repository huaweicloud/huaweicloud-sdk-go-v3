package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// QuotaInfo 配额信息，包括总配额和已使用配额。
type QuotaInfo struct {

	// 资源类型。  取值：loadbalancer、listener、ipgroup、pool、member、members_per_pool、 healthmonitor、l7policy、certificate、security_policy、condition_per_policy、listeners_per_pool、ipgroup_bindings、ipgroup_bindings，  其中members_per_pool表示一个pool下最多可关联的member数量。
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
