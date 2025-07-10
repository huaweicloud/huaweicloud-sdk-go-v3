package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// AdInfo 域信息。
type AdInfo struct {

	// 域类型。 - LITE_AS：本地认证。 - LOCAL_AD：本地AD。
	DomainType *AdInfoDomainType `json:"domain_type,omitempty"`

	// 域名，域类型为LOCAL_AD时有值。
	DomainName *string `json:"domain_name,omitempty"`

	// 域管理员账号，域类型为LOCAL_AD时有值。
	DomainAdminAccount *string `json:"domain_admin_account,omitempty"`

	// 主域控制器名称，域类型为LOCAL_AD时有值。
	ActiveDomainName *string `json:"active_domain_name,omitempty"`

	// 主域控制器IP地址，域类型为LOCAL_AD时有值。
	ActiveDomainIp *string `json:"active_domain_ip,omitempty"`

	// 备域控制器名称，域类型为LOCAL_AD时有值。
	StandbyDomainName *string `json:"standby_domain_name,omitempty"`

	// 备域控制器IP地址，域类型为LOCAL_AD时有值。
	StandbyDomainIp *string `json:"standby_domain_ip,omitempty"`

	// 主DNS IP地址，域类型为LOCAL_AD时有值。
	ActiveDnsIp *string `json:"active_dns_ip,omitempty"`

	// 备DNS IP地址，域类型为LOCAL_AD时有值。
	StandbyDnsIp *string `json:"standby_dns_ip,omitempty"`

	// 是否在删除桌面的同时删除AD上对应的计算机对象，'0'代表不删除，'1'代表删除。
	DeleteComputerObject *string `json:"delete_computer_object,omitempty"`

	// 是否开启LDAPS。
	UseLdaps *bool `json:"use_ldaps,omitempty"`

	TlsConfig *TlsConfig `json:"tls_config,omitempty"`
}

func (o AdInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AdInfo struct{}"
	}

	return strings.Join([]string{"AdInfo", string(data)}, " ")
}

type AdInfoDomainType struct {
	value string
}

type AdInfoDomainTypeEnum struct {
	LITE_AS  AdInfoDomainType
	LOCAL_AD AdInfoDomainType
}

func GetAdInfoDomainTypeEnum() AdInfoDomainTypeEnum {
	return AdInfoDomainTypeEnum{
		LITE_AS: AdInfoDomainType{
			value: "LITE_AS",
		},
		LOCAL_AD: AdInfoDomainType{
			value: "LOCAL_AD",
		},
	}
}

func (c AdInfoDomainType) Value() string {
	return c.value
}

func (c AdInfoDomainType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AdInfoDomainType) UnmarshalJSON(b []byte) error {
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
