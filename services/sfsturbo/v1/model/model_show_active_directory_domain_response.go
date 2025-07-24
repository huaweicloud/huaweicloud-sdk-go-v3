package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowActiveDirectoryDomainResponse Response Object
type ShowActiveDirectoryDomainResponse struct {

	// 域控服务器的域名，在创建域服务器时指定
	DomainName *string `json:"domain_name,omitempty"`

	// 存储系统在AD域中的名称
	SystemName *string `json:"system_name,omitempty"`

	// DNS服务器IP地址，用于解析AD域的域名
	DnsServer *[]string `json:"dns_server,omitempty"`

	// 域中包含的一类目录对象如用户、计算机、打印机等资源的总称，加入后将作为其中的一员。若不填，则默认加入到computers组织单元。
	OrganizationUnit *string `json:"organization_unit,omitempty"`

	// vpc的id
	VpcId *string `json:"vpc_id,omitempty"`

	// AD域当前状态信息
	Status *ShowActiveDirectoryDomainResponseStatus `json:"status,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowActiveDirectoryDomainResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowActiveDirectoryDomainResponse struct{}"
	}

	return strings.Join([]string{"ShowActiveDirectoryDomainResponse", string(data)}, " ")
}

type ShowActiveDirectoryDomainResponseStatus struct {
	value string
}

type ShowActiveDirectoryDomainResponseStatusEnum struct {
	JOINING   ShowActiveDirectoryDomainResponseStatus
	AVAILABLE ShowActiveDirectoryDomainResponseStatus
	EXITING   ShowActiveDirectoryDomainResponseStatus
	FAILED    ShowActiveDirectoryDomainResponseStatus
}

func GetShowActiveDirectoryDomainResponseStatusEnum() ShowActiveDirectoryDomainResponseStatusEnum {
	return ShowActiveDirectoryDomainResponseStatusEnum{
		JOINING: ShowActiveDirectoryDomainResponseStatus{
			value: "JOINING",
		},
		AVAILABLE: ShowActiveDirectoryDomainResponseStatus{
			value: "AVAILABLE",
		},
		EXITING: ShowActiveDirectoryDomainResponseStatus{
			value: "EXITING",
		},
		FAILED: ShowActiveDirectoryDomainResponseStatus{
			value: "FAILED",
		},
	}
}

func (c ShowActiveDirectoryDomainResponseStatus) Value() string {
	return c.value
}

func (c ShowActiveDirectoryDomainResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowActiveDirectoryDomainResponseStatus) UnmarshalJSON(b []byte) error {
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
