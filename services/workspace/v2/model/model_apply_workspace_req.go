package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ApplyWorkspaceReq 申请开通云办公服务请求
type ApplyWorkspaceReq struct {

	// 主认证方式。 - KERBEROS：KERBEROS。 - KERBEROS_THIRD_SSO：第三方登录认证。
	AuthType *ApplyWorkspaceReqAuthType `json:"auth_type,omitempty"`

	AdDomains *AdDomain `json:"ad_domains"`

	ThirdGatewayInfo *ThirdGatewayConfigInfo `json:"third_gateway_info,omitempty"`

	// 企业ID。 企业ID是您在云桌面服务的唯一标识，终端用户登录时需要填写企业ID，若不自定义设置企业ID，系统将自动生成您的企业ID。格式为由半角数字、字母、_-组成，长度范围小于等于32个字符。
	EnterpriseId *string `json:"enterprise_id,omitempty"`

	// VPC ID。
	VpcId *string `json:"vpc_id,omitempty"`

	// 指定业务子网的网络ID，子网不能与172.16.0.0/12冲突。
	SubnetIds *[]Subnet `json:"subnet_ids,omitempty"`

	// 管理子网网段。 注：不能与172.16.0.0/12以及subnet_ids参数下子网的网段冲突。
	ManageSubnetCidr *string `json:"manage_subnet_cidr,omitempty"`

	// 接入方式。 - INTERNET：表示Internet接入。 - DEDICATED：表示专线接入。 - BOTH：表示两种接入方式都支持。
	AccessMode *ApplyWorkspaceReqAccessMode `json:"access_mode,omitempty"`

	ApplySharedVpcDedicatedParam *ApplySharedVpcDedicatedParam `json:"apply_shared_vpc_dedicated_param,omitempty"`

	// 专线接入网段列表，多个网段信息用分号隔开，列表长度不超过5。
	DedicatedSubnets *string `json:"dedicated_subnets,omitempty"`

	// 开通服务资源使用的可用分区。
	AvailabilityZone *string `json:"availability_zone,omitempty"`

	// 外部网络。
	PublicipType *string `json:"publicip_type,omitempty"`

	AssistAuthConfig *AssistAuthMethodConfigRequest `json:"assist_auth_config,omitempty"`

	// 边缘开户信息
	SiteConfigs *[]SiteConfigsRequest `json:"site_configs,omitempty"`

	// 桌面退订是否发送邮件通知。
	IsSendEmail *bool `json:"is_send_email,omitempty"`

	// 企业项目ID，默认\"0\"
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
}

func (o ApplyWorkspaceReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApplyWorkspaceReq struct{}"
	}

	return strings.Join([]string{"ApplyWorkspaceReq", string(data)}, " ")
}

type ApplyWorkspaceReqAuthType struct {
	value string
}

type ApplyWorkspaceReqAuthTypeEnum struct {
	KERBEROS           ApplyWorkspaceReqAuthType
	KERBEROS_THIRD_SSO ApplyWorkspaceReqAuthType
}

func GetApplyWorkspaceReqAuthTypeEnum() ApplyWorkspaceReqAuthTypeEnum {
	return ApplyWorkspaceReqAuthTypeEnum{
		KERBEROS: ApplyWorkspaceReqAuthType{
			value: "KERBEROS",
		},
		KERBEROS_THIRD_SSO: ApplyWorkspaceReqAuthType{
			value: "KERBEROS_THIRD_SSO",
		},
	}
}

func (c ApplyWorkspaceReqAuthType) Value() string {
	return c.value
}

func (c ApplyWorkspaceReqAuthType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ApplyWorkspaceReqAuthType) UnmarshalJSON(b []byte) error {
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

type ApplyWorkspaceReqAccessMode struct {
	value string
}

type ApplyWorkspaceReqAccessModeEnum struct {
	INTERNET  ApplyWorkspaceReqAccessMode
	DEDICATED ApplyWorkspaceReqAccessMode
	BOTH      ApplyWorkspaceReqAccessMode
}

func GetApplyWorkspaceReqAccessModeEnum() ApplyWorkspaceReqAccessModeEnum {
	return ApplyWorkspaceReqAccessModeEnum{
		INTERNET: ApplyWorkspaceReqAccessMode{
			value: "INTERNET",
		},
		DEDICATED: ApplyWorkspaceReqAccessMode{
			value: "DEDICATED",
		},
		BOTH: ApplyWorkspaceReqAccessMode{
			value: "BOTH",
		},
	}
}

func (c ApplyWorkspaceReqAccessMode) Value() string {
	return c.value
}

func (c ApplyWorkspaceReqAccessMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ApplyWorkspaceReqAccessMode) UnmarshalJSON(b []byte) error {
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
