package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// FindingDetails 访问分析结果的详细信息。
type FindingDetails struct {
	ExternalAccessDetails *ExternalAccessDetails `json:"external_access_details,omitempty"`

	PrivilegeEscalationDetails *PrivilegeEscalationDetails `json:"privilege_escalation_details,omitempty"`

	UnusedIamUserAccessKeyDetails *UnusedIamUserAccessKeyDetails `json:"unused_iam_user_access_key_details,omitempty"`

	UnusedIamUserPasswordDetails *UnusedIamUserPasswordDetails `json:"unused_iam_user_password_details,omitempty"`

	UnusedPermissionDetails *UnusedPermissionDetails `json:"unused_permission_details,omitempty"`

	UnusedIamAgencyDetails *UnusedIamAgencyDetails `json:"unused_iam_agency_details,omitempty"`

	IamBpRootUserHasAccessKeyDetails *IamBpRootUserHasAccessKeyDetails `json:"iam_bp_root_user_has_access_key_details,omitempty"`

	IamBpAccessApiWithPasswordDetails *IamBpAccessApiWithPasswordDetails `json:"iam_bp_access_api_with_password_details,omitempty"`

	IamBpLoginProtectionDisabledDetails *IamBpLoginProtectionDisabledDetails `json:"iam_bp_login_protection_disabled_details,omitempty"`

	IamBpMfaUnconfiguredDetails *IamBpMfaUnconfiguredDetails `json:"iam_bp_mfa_unconfigured_details,omitempty"`

	IamBpAssignHighRiskSysPolicyOrRoleToUserDetails *IamBpAssignHighRiskSysPolicyOrRoleToUserDetails `json:"iam_bp_assign_high_risk_sys_policy_or_role_to_user_details,omitempty"`

	IamBpAttachHighRiskSysIdentityPolicyToUserDetails *IamBpAttachHighRiskSysIdentityPolicyToUserDetails `json:"iam_bp_attach_high_risk_sys_identity_policy_to_user_details,omitempty"`

	IamBpAssignHighRiskSysPolicyOrRoleToAgencyDetails *IamBpAssignHighRiskSysPolicyOrRoleToAgencyDetails `json:"iam_bp_assign_high_risk_sys_policy_or_role_to_agency_details,omitempty"`

	IamBpAttachHighRiskSysIdentityPolicyToAgencyDetails *IamBpAttachHighRiskSysIdentityPolicyToAgencyDetails `json:"iam_bp_attach_high_risk_sys_identity_policy_to_agency_details,omitempty"`
}

func (o FindingDetails) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FindingDetails struct{}"
	}

	return strings.Join([]string{"FindingDetails", string(data)}, " ")
}
