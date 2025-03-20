package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GetAccountSummaryV5Response Response Object
type GetAccountSummaryV5Response struct {

	// 附加到委托或信任委托上的身份策略的最大数量。
	AttachedPoliciesPerAgencyQuota *int32 `json:"attached_policies_per_agency_quota,omitempty"`

	// 附加到用户组上的身份策略的最大数量。
	AttachedPoliciesPerGroupQuota *int32 `json:"attached_policies_per_group_quota,omitempty"`

	// 附加到IAM用户上的身份策略的最大数量。
	AttachedPoliciesPerUserQuota *int32 `json:"attached_policies_per_user_quota,omitempty"`

	// 自定义身份策略的最大数量。
	PoliciesQuota *int32 `json:"policies_quota,omitempty"`

	// 身份策略及信任策略的策略文档的最大字符数，不包括空格。
	PolicySizeQuota *int32 `json:"policy_size_quota,omitempty"`

	// 自定义身份策略同一时刻保留的最大版本数量。
	VersionsPerPolicyQuota *int32 `json:"versions_per_policy_quota,omitempty"`

	// 此账号中当前创建的自定义身份策略数量。
	Policies *int32 `json:"policies,omitempty"`

	// 此账号中当前创建的委托及信任委托的总数量。
	Agencies *int32 `json:"agencies,omitempty"`

	// 此账号能够创建的委托及信任委托的总数上限。
	AgenciesQuota *int32 `json:"agencies_quota,omitempty"`

	// 此账号当前创建的IAM用户数量，包括根用户。
	Users *int32 `json:"users,omitempty"`

	// 此账号能够创建的IAM用户数上限，包括根用户。
	UsersQuota *int32 `json:"users_quota,omitempty"`

	// 此账号当前创建的用户组数量。
	Groups *int32 `json:"groups,omitempty"`

	// 此账号能够创建的用户组数上限。
	GroupsQuota *int32 `json:"groups_quota,omitempty"`

	// 根用户绑定的已启用MFA的数量。
	RootUserMfaEnabled *int32 `json:"root_user_mfa_enabled,omitempty"`
	HttpStatusCode     int    `json:"-"`
}

func (o GetAccountSummaryV5Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetAccountSummaryV5Response struct{}"
	}

	return strings.Join([]string{"GetAccountSummaryV5Response", string(data)}, " ")
}
