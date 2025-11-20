package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type FindingFilter struct {
	Criterion *Criterion `json:"criterion"`

	// 过滤键。 - resource：资源URN - resource_type：资源类型 - resource_owner_account：资源所有者账号 - is_public：公共访问权限 - id：分析结果ID - status：分析结果类型 - principal_type：主体类型 - principal_identifier：主体Identifier - change_type：分析结果状态的变化 - existing_finding_id：已有分析结果ID - existing_finding_status：已有分析结果状态 - condition.g:PrincipalUrn：主体URN - condition.g:PrincipalId：主体ID - condition.g:PrincipalAccount：主体账号 - condition.g:PrincipalOrgId：主体OrgID - condition.g:PrincipalOrgPath：主体组织路径 - condition.g:PrincipalOrgManagementAccountId：主体组织管理账号ID - condition.g:SourceIp：源IP - condition.g:SourceVpc：源VPC - condition.g:SourceVpce：源VPCE - finding_type：分析结果类型
	Key FindingFilterKey `json:"key"`
}

func (o FindingFilter) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FindingFilter struct{}"
	}

	return strings.Join([]string{"FindingFilter", string(data)}, " ")
}

type FindingFilterKey struct {
	value string
}

type FindingFilterKeyEnum struct {
	RESOURCE                                       FindingFilterKey
	RESOURCE_TYPE                                  FindingFilterKey
	RESOURCE_OWNER_ACCOUNT                         FindingFilterKey
	IS_PUBLIC                                      FindingFilterKey
	ID                                             FindingFilterKey
	STATUS                                         FindingFilterKey
	PRINCIPAL_TYPE                                 FindingFilterKey
	PRINCIPAL_IDENTIFIER                           FindingFilterKey
	CHANGE_TYPE                                    FindingFilterKey
	EXISTING_FINDING_ID                            FindingFilterKey
	EXISTING_FINDING_STATUS                        FindingFilterKey
	CONDITION_GPRINCIPAL_URN                       FindingFilterKey
	CONDITION_GPRINCIPAL_ID                        FindingFilterKey
	CONDITION_GPRINCIPAL_ACCOUNT                   FindingFilterKey
	CONDITION_GPRINCIPAL_ORG_ID                    FindingFilterKey
	CONDITION_GPRINCIPAL_ORG_PATH                  FindingFilterKey
	CONDITION_GPRINCIPAL_ORG_MANAGEMENT_ACCOUNT_ID FindingFilterKey
	CONDITION_GSOURCE_IP                           FindingFilterKey
	CONDITION_GSOURCE_VPC                          FindingFilterKey
	CONDITION_GSOURCE_VPCE                         FindingFilterKey
	FINDING_TYPE                                   FindingFilterKey
}

func GetFindingFilterKeyEnum() FindingFilterKeyEnum {
	return FindingFilterKeyEnum{
		RESOURCE: FindingFilterKey{
			value: "resource",
		},
		RESOURCE_TYPE: FindingFilterKey{
			value: "resource_type",
		},
		RESOURCE_OWNER_ACCOUNT: FindingFilterKey{
			value: "resource_owner_account",
		},
		IS_PUBLIC: FindingFilterKey{
			value: "is_public",
		},
		ID: FindingFilterKey{
			value: "id",
		},
		STATUS: FindingFilterKey{
			value: "status",
		},
		PRINCIPAL_TYPE: FindingFilterKey{
			value: "principal_type",
		},
		PRINCIPAL_IDENTIFIER: FindingFilterKey{
			value: "principal_identifier",
		},
		CHANGE_TYPE: FindingFilterKey{
			value: "change_type",
		},
		EXISTING_FINDING_ID: FindingFilterKey{
			value: "existing_finding_id",
		},
		EXISTING_FINDING_STATUS: FindingFilterKey{
			value: "existing_finding_status",
		},
		CONDITION_GPRINCIPAL_URN: FindingFilterKey{
			value: "condition.g:PrincipalUrn",
		},
		CONDITION_GPRINCIPAL_ID: FindingFilterKey{
			value: "condition.g:PrincipalId",
		},
		CONDITION_GPRINCIPAL_ACCOUNT: FindingFilterKey{
			value: "condition.g:PrincipalAccount",
		},
		CONDITION_GPRINCIPAL_ORG_ID: FindingFilterKey{
			value: "condition.g:PrincipalOrgId",
		},
		CONDITION_GPRINCIPAL_ORG_PATH: FindingFilterKey{
			value: "condition.g:PrincipalOrgPath",
		},
		CONDITION_GPRINCIPAL_ORG_MANAGEMENT_ACCOUNT_ID: FindingFilterKey{
			value: "condition.g:PrincipalOrgManagementAccountId",
		},
		CONDITION_GSOURCE_IP: FindingFilterKey{
			value: "condition.g:SourceIp",
		},
		CONDITION_GSOURCE_VPC: FindingFilterKey{
			value: "condition.g:SourceVpc",
		},
		CONDITION_GSOURCE_VPCE: FindingFilterKey{
			value: "condition.g:SourceVpce",
		},
		FINDING_TYPE: FindingFilterKey{
			value: "finding_type",
		},
	}
}

func (c FindingFilterKey) Value() string {
	return c.value
}

func (c FindingFilterKey) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *FindingFilterKey) UnmarshalJSON(b []byte) error {
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
