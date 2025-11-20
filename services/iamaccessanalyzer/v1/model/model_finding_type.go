package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// FindingType 访问分析结果类型。 - external_access：外部访问 - privilege_escalation：提权访问 - unused_iam_user_access_key：未使用访问密钥 - unused_iam_user_password：未使用密码 - unused_permission：未使用权限 - unused_iam_agency：未使用委托 - iam_bp_root_user_has_access_key：为根用户绑定AK/SK - iam_bp_access_api_with_password：使用密码访问API - iam_bp_login_protection_disabled：未开启登录保护 - iam_bp_mfa_unconfigured：未绑定MFA - iam_bp_assign_high_risk_sys_policy_or_role_to_user：为用户授予高风险系统策略或角色 - iam_bp_attach_high_risk_sys_identity_policy_to_user：为用户授予高风险系统身份策略 - iam_bp_assign_high_risk_sys_policy_or_role_to_agency：为委托授予高风险系统策略或角色 - iam_bp_attach_high_risk_sys_identity_policy_to_agency：为委托授予高风险系统身份策略
type FindingType struct {
	value string
}

type FindingTypeEnum struct {
	EXTERNAL_ACCESS                                       FindingType
	PRIVILEGE_ESCALATION                                  FindingType
	UNUSED_IAM_USER_ACCESS_KEY                            FindingType
	UNUSED_IAM_USER_PASSWORD                              FindingType
	UNUSED_PERMISSION                                     FindingType
	UNUSED_IAM_AGENCY                                     FindingType
	IAM_BP_ROOT_USER_HAS_ACCESS_KEY                       FindingType
	IAM_BP_ACCESS_API_WITH_PASSWORD                       FindingType
	IAM_BP_LOGIN_PROTECTION_DISABLED                      FindingType
	IAM_BP_MFA_UNCONFIGURED                               FindingType
	IAM_BP_ASSIGN_HIGH_RISK_SYS_POLICY_OR_ROLE_TO_USER    FindingType
	IAM_BP_ATTACH_HIGH_RISK_SYS_IDENTITY_POLICY_TO_USER   FindingType
	IAM_BP_ASSIGN_HIGH_RISK_SYS_POLICY_OR_ROLE_TO_AGENCY  FindingType
	IAM_BP_ATTACH_HIGH_RISK_SYS_IDENTITY_POLICY_TO_AGENCY FindingType
}

func GetFindingTypeEnum() FindingTypeEnum {
	return FindingTypeEnum{
		EXTERNAL_ACCESS: FindingType{
			value: "external_access",
		},
		PRIVILEGE_ESCALATION: FindingType{
			value: "privilege_escalation",
		},
		UNUSED_IAM_USER_ACCESS_KEY: FindingType{
			value: "unused_iam_user_access_key",
		},
		UNUSED_IAM_USER_PASSWORD: FindingType{
			value: "unused_iam_user_password",
		},
		UNUSED_PERMISSION: FindingType{
			value: "unused_permission",
		},
		UNUSED_IAM_AGENCY: FindingType{
			value: "unused_iam_agency",
		},
		IAM_BP_ROOT_USER_HAS_ACCESS_KEY: FindingType{
			value: "iam_bp_root_user_has_access_key",
		},
		IAM_BP_ACCESS_API_WITH_PASSWORD: FindingType{
			value: "iam_bp_access_api_with_password",
		},
		IAM_BP_LOGIN_PROTECTION_DISABLED: FindingType{
			value: "iam_bp_login_protection_disabled",
		},
		IAM_BP_MFA_UNCONFIGURED: FindingType{
			value: "iam_bp_mfa_unconfigured",
		},
		IAM_BP_ASSIGN_HIGH_RISK_SYS_POLICY_OR_ROLE_TO_USER: FindingType{
			value: "iam_bp_assign_high_risk_sys_policy_or_role_to_user",
		},
		IAM_BP_ATTACH_HIGH_RISK_SYS_IDENTITY_POLICY_TO_USER: FindingType{
			value: "iam_bp_attach_high_risk_sys_identity_policy_to_user",
		},
		IAM_BP_ASSIGN_HIGH_RISK_SYS_POLICY_OR_ROLE_TO_AGENCY: FindingType{
			value: "iam_bp_assign_high_risk_sys_policy_or_role_to_agency",
		},
		IAM_BP_ATTACH_HIGH_RISK_SYS_IDENTITY_POLICY_TO_AGENCY: FindingType{
			value: "iam_bp_attach_high_risk_sys_identity_policy_to_agency",
		},
	}
}

func (c FindingType) Value() string {
	return c.value
}

func (c FindingType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *FindingType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}
