package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ApplyWorkspaceAdDomain 域信息。
type ApplyWorkspaceAdDomain struct {

	// 域类型。 - LITE_AS：本地认证。 - LOCAL_AD：本地AD。 说明：域类型为“LOCAL_AD”时，请确保所选VPC网络与AD所属网络可连通。
	DomainType ApplyWorkspaceAdDomainDomainType `json:"domain_type"`

	// 域名称。域类型为LOCAL_AD时需要配置。 域名必须为AD服务器上已经存在的域名，且长度不超过55。
	DomainName *string `json:"domain_name,omitempty"`

	// 域管理员账号。域类型为LOCAL_AD时需要配置。 必须为AD服务器上已经存在的域管理员账号。
	DomainAdminAccount *string `json:"domain_admin_account,omitempty"`

	// 域管理员账号密码。域类型为LOCAL_AD时需要配置。
	DomainPassword *string `json:"domain_password,omitempty"`

	// 主域控制器IP地址。域类型为LOCAL_AD时需要配置。
	ActiveDomainIp *string `json:"active_domain_ip,omitempty"`

	// 主域控制器名称。域类型为LOCAL_AD时需要配置。
	ActiveDomainName *string `json:"active_domain_name,omitempty"`

	// 备域控制器IP地址。域类型为LOCAL_AD时且配置备节点时需要配置。
	StandbyDomainIp *string `json:"standby_domain_ip,omitempty"`

	// 备域控制器名称。域类型为LOCAL_AD时且配置备节点时需要配置。
	StandbyDomainName *string `json:"standby_domain_name,omitempty"`

	// 主DNS IP地址。域类型为LOCAL_AD时需要配置。
	ActiveDnsIp *string `json:"active_dns_ip,omitempty"`

	// 备DNS IP地址。域类型为LOCAL_AD时且配置备节点时需要配置。
	StandbyDnsIp *string `json:"standby_dns_ip,omitempty"`

	// 是否在删除桌面的同时删除AD上对应的计算机对象，0代表不删除，1代表删除。
	DeleteComputerObject *int32 `json:"delete_computer_object,omitempty"`

	// 是否开启LDAPS。
	UseLdaps *bool `json:"use_ldaps,omitempty"`

	TlsConfig *TlsConfig `json:"tls_config,omitempty"`
}

func (o ApplyWorkspaceAdDomain) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApplyWorkspaceAdDomain struct{}"
	}

	return strings.Join([]string{"ApplyWorkspaceAdDomain", string(data)}, " ")
}

type ApplyWorkspaceAdDomainDomainType struct {
	value string
}

type ApplyWorkspaceAdDomainDomainTypeEnum struct {
	LITE_AS  ApplyWorkspaceAdDomainDomainType
	LOCAL_AD ApplyWorkspaceAdDomainDomainType
}

func GetApplyWorkspaceAdDomainDomainTypeEnum() ApplyWorkspaceAdDomainDomainTypeEnum {
	return ApplyWorkspaceAdDomainDomainTypeEnum{
		LITE_AS: ApplyWorkspaceAdDomainDomainType{
			value: "LITE_AS",
		},
		LOCAL_AD: ApplyWorkspaceAdDomainDomainType{
			value: "LOCAL_AD",
		},
	}
}

func (c ApplyWorkspaceAdDomainDomainType) Value() string {
	return c.value
}

func (c ApplyWorkspaceAdDomainDomainType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ApplyWorkspaceAdDomainDomainType) UnmarshalJSON(b []byte) error {
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
