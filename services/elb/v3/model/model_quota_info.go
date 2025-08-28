package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// QuotaInfo **参数解释**：配额信息，包括总配额和已使用配额。
type QuotaInfo struct {

	// **参数解释**：配额类型。  **取值范围**：loadbalancer、listener、ipgroup、pool、member、healthmonitor、l7policy、certificate、security_policy、listeners_per_loadbalancer、listeners_per_pool、members_per_pool、condition_per_policy、ipgroup_bindings、ipgroup_max_length、ipgroups_per_listener、pools_per_l7policy、l7policies_per_listener、free_instance_members_per_pool、free_instance_listeners_per_loadbalancer。
	QuotaKey string `json:"quota_key"`

	// **参数解释**：总配额。  **取值范围**： - 大于等于0：表示当前配额数量。 - -1：表示无配额限制。
	QuotaLimit int32 `json:"quota_limit"`

	// **参数解释**：已使用配额。  **取值范围**：大于等于0。
	Used int32 `json:"used"`

	// **参数解释**：配额单位。  **取值范围**：count，表示个数。
	Unit string `json:"unit"`
}

func (o QuotaInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QuotaInfo struct{}"
	}

	return strings.Join([]string{"QuotaInfo", string(data)}, " ")
}
