package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DomainStatuses struct {

	// **参数解释**： 服务名。 **约束限制**： 不涉及。 **取值范围**： 默认CodeCI。
	ServiceName *string `json:"service_name,omitempty"`

	// **参数解释**： 租户id。 **约束限制**： 不涉及。 **取值范围**： 不涉及。
	DomainId *string `json:"domain_id,omitempty"`

	// **参数解释**： 租户名称。 **约束限制**： 不涉及。 **取值范围**： 不涉及。
	DomainName *string `json:"domain_name,omitempty"`

	// **参数解释**： freeQuota。 **约束限制**： 不涉及。 **取值范围**： true或false。
	FreeQuota *bool `json:"free_quota,omitempty"`

	// **参数解释**： freePackageQuota。 **约束限制**： 不涉及。 **取值范围**： true或false。
	FreePackageQuota *bool `json:"free_package_quota,omitempty"`

	// **参数解释**： 状态。 **约束限制**： 不涉及。 **取值范围**： 不涉及。
	Status *int32 `json:"status,omitempty"`

	// **参数解释**： isFederation。 **约束限制**： 不涉及。 **取值范围**： 不涉及。
	IsFederation *int32 `json:"is_federation,omitempty"`

	// **参数解释**： 白名单标识。 **约束限制**： 不涉及。 **取值范围**： 不涉及。
	IsWhitelist *int32 `json:"is_whitelist,omitempty"`

	// **参数解释**： 地区。 **约束限制**： 不涉及。 **取值范围**： 不涉及。
	Region *string `json:"region,omitempty"`

	// **参数解释**： 包类型。 **约束限制**： 不涉及。 **取值范围**： 不涉及。
	PackageType *string `json:"package_type,omitempty"`

	// **参数解释**： 租户类型。 **约束限制**： 不涉及。 **取值范围**： 不涉及。
	DomainType *int32 `json:"domain_type,omitempty"`

	// **参数解释**： 租户业务类型。 **约束限制**： 不涉及。 **取值范围**： 不涉及。
	DomainBusinessType *string `json:"domain_business_type,omitempty"`

	// **参数解释**： domainMaliciousDockerArgs。 **约束限制**： 不涉及。 **取值范围**： 不涉及。
	DomainMaliciousDockerArgs *string `json:"domain_malicious_docker_args,omitempty"`

	// **参数解释**： 允许自定义环境数。 **约束限制**： 不涉及。 **取值范围**： 不涉及。
	AllowCustomEnv *int32 `json:"allow_custom_env,omitempty"`

	// **参数解释**： 默认的加速阈值注入。 **约束限制**： 不涉及。 **取值范围**： 不涉及。
	SpecThreshold *string `json:"spec_threshold,omitempty"`
}

func (o DomainStatuses) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DomainStatuses struct{}"
	}

	return strings.Join([]string{"DomainStatuses", string(data)}, " ")
}
