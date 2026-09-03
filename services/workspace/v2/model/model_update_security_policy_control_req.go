package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateSecurityPolicyControlReq Update security policy control request.
type UpdateSecurityPolicyControlReq struct {

	// 需要开启安全策略管控的实例 ID 列表。
	EnabledInstances *[]string `json:"enabled_instances,omitempty"`

	// 需要开启安全策略管控的标签列表，格式为 key:value。
	EnabledTags *[]string `json:"enabled_tags,omitempty"`

	// 需要关闭安全策略管控的实例 ID 列表。
	DisabledInstances *[]string `json:"disabled_instances,omitempty"`

	// 需要关闭安全策略管控的标签列表，格式为 key:value。
	DisabledTags *[]string `json:"disabled_tags,omitempty"`
}

func (o UpdateSecurityPolicyControlReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSecurityPolicyControlReq struct{}"
	}

	return strings.Join([]string{"UpdateSecurityPolicyControlReq", string(data)}, " ")
}
