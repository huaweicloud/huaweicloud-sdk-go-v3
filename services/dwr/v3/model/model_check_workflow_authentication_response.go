package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Response Object
type CheckWorkflowAuthenticationResponse struct {

	// 创建时间。
	CreateTime *string `json:"create_time,omitempty"`

	// 最近修改时间。
	LastModifyTime *string `json:"last_modify_time,omitempty"`

	// 委托方帐号ID。
	AgencyId *string `json:"agency_id,omitempty"`

	// 委托名。
	AgencyName *string `json:"agency_name,omitempty"`

	// 委托的期限。取值为\"FOREVER\"或“null”表示委托的期限为永久，取值为\"ONEDAY\"表示委托的期限为一天。
	AgencyDuration *CheckWorkflowAuthenticationResponseAgencyDuration `json:"agency_duration,omitempty"`

	// 被委托方帐号名。
	TrustDomainName *string `json:"trust_domain_name,omitempty"`

	// 权限ID。
	RoleId *string `json:"role_id,omitempty"`

	// 权限使用的依赖函数。
	RoleDependentByFunction *string `json:"role_dependent_by_function,omitempty"`

	// 权限备注名。
	RoleRemarkName *string `json:"role_remark_name,omitempty"`

	// 权限的备注模式： AX表示在domain层显示。 XA表示在project层显示。 AA表示在domain和project层均显示。 XX表示在domain和project层均不显示。 自定义策略的显示模式只能为AX或者XA，不能在domain层和project层都显示（AA），或者在domain层和project层都不显示（XX）
	RoleRemarkType *CheckWorkflowAuthenticationResponseRoleRemarkType `json:"role_remark_type,omitempty"`

	XRequestId *string `json:"x-request-id,omitempty"`

	Connection *string `json:"Connection,omitempty"`

	ContentLength *string `json:"Content-Length,omitempty"`

	Date           *string `json:"Date,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CheckWorkflowAuthenticationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckWorkflowAuthenticationResponse struct{}"
	}

	return strings.Join([]string{"CheckWorkflowAuthenticationResponse", string(data)}, " ")
}

type CheckWorkflowAuthenticationResponseAgencyDuration struct {
	value string
}

type CheckWorkflowAuthenticationResponseAgencyDurationEnum struct {
	FOREVER CheckWorkflowAuthenticationResponseAgencyDuration
	ONEDAY  CheckWorkflowAuthenticationResponseAgencyDuration
}

func GetCheckWorkflowAuthenticationResponseAgencyDurationEnum() CheckWorkflowAuthenticationResponseAgencyDurationEnum {
	return CheckWorkflowAuthenticationResponseAgencyDurationEnum{
		FOREVER: CheckWorkflowAuthenticationResponseAgencyDuration{
			value: "FOREVER",
		},
		ONEDAY: CheckWorkflowAuthenticationResponseAgencyDuration{
			value: "ONEDAY",
		},
	}
}

func (c CheckWorkflowAuthenticationResponseAgencyDuration) Value() string {
	return c.value
}

func (c CheckWorkflowAuthenticationResponseAgencyDuration) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CheckWorkflowAuthenticationResponseAgencyDuration) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}

type CheckWorkflowAuthenticationResponseRoleRemarkType struct {
	value string
}

type CheckWorkflowAuthenticationResponseRoleRemarkTypeEnum struct {
	AX CheckWorkflowAuthenticationResponseRoleRemarkType
	XA CheckWorkflowAuthenticationResponseRoleRemarkType
	AA CheckWorkflowAuthenticationResponseRoleRemarkType
	XX CheckWorkflowAuthenticationResponseRoleRemarkType
}

func GetCheckWorkflowAuthenticationResponseRoleRemarkTypeEnum() CheckWorkflowAuthenticationResponseRoleRemarkTypeEnum {
	return CheckWorkflowAuthenticationResponseRoleRemarkTypeEnum{
		AX: CheckWorkflowAuthenticationResponseRoleRemarkType{
			value: "AX",
		},
		XA: CheckWorkflowAuthenticationResponseRoleRemarkType{
			value: "XA",
		},
		AA: CheckWorkflowAuthenticationResponseRoleRemarkType{
			value: "AA",
		},
		XX: CheckWorkflowAuthenticationResponseRoleRemarkType{
			value: "XX",
		},
	}
}

func (c CheckWorkflowAuthenticationResponseRoleRemarkType) Value() string {
	return c.value
}

func (c CheckWorkflowAuthenticationResponseRoleRemarkType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CheckWorkflowAuthenticationResponseRoleRemarkType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
