package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateActiveDirectoryDomainRequestBody AD域配置信息
type UpdateActiveDirectoryDomainRequestBody struct {

	// 服务账号，在创建域服务器时指定，一般默认为administrator
	ServiceAccount string `json:"service_account"`

	// 账号对应密码
	Password string `json:"password"`

	// 域控服务器的域名，在创建域服务器时指定
	DomainName string `json:"domain_name"`

	// 存储系统在AD域中的名称
	SystemName string `json:"system_name"`

	// 如果域控制器中已存在同系统名称的存储系统，开启该选项后，将覆盖原有的存储系统信息。
	OverwriteSameAccount *bool `json:"overwrite_same_account,omitempty"`

	// DNS服务器IP地址，用于解析AD域的域名
	DnsServer []string `json:"dns_server"`

	// 域中包含的一类目录对象如用户、计算机、打印机等资源的总称，加入后将作为其中的一员。若不填，则默认加入到computers组织单元。
	OrganizationUnit *string `json:"organization_unit,omitempty"`

	// vpc的id
	VpcId *string `json:"vpc_id,omitempty"`
}

func (o UpdateActiveDirectoryDomainRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateActiveDirectoryDomainRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateActiveDirectoryDomainRequestBody", string(data)}, " ")
}
