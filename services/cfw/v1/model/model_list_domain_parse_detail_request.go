package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDomainParseDetailRequest Request Object
type ListDomainParseDetailRequest struct {

	// 域名，如www.test.com
	DomainName string `json:"domain_name"`

	// 企业项目ID，用户根据组织规划企业项目，对应的ID为企业项目ID，可通过[如何获取企业项目ID](cfw_02_0027.xml)获取，用户未开启企业项目时为0
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 防火墙id，可通过[防火墙ID获取方式](cfw_02_0028.xml)获取
	FwInstanceId *string `json:"fw_instance_id,omitempty"`

	// 地址类型，0 ipv4,1 ipv6
	AddressType *int32 `json:"address_type,omitempty"`
}

func (o ListDomainParseDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDomainParseDetailRequest struct{}"
	}

	return strings.Join([]string{"ListDomainParseDetailRequest", string(data)}, " ")
}
