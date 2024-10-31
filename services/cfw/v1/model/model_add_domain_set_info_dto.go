package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AddDomainSetInfoDto struct {

	// 防火墙id，可通过[防火墙ID获取方式](cfw_02_0028.xml)获取
	FwInstanceId *string `json:"fw_instance_id,omitempty"`

	// 防护对象id，是创建云防火墙后用于区分互联网边界防护和VPC边界防护的标志id，可通过调用[查询防火墙实例接口](ListFirewallDetail.xml)获得，通过返回值中的data.records.protect_objects.object_id（.表示各对象之间层级的区分）获得，注意type为0的为互联网边界防护对象id，type为1的为VPC边界防护对象id，type可通过data.records.protect_objects.type（.表示各对象之间层级的区分）获得
	ObjectId string `json:"object_id"`

	// 域名组名称
	Name string `json:"name"`

	// 域名组描述
	Description *string `json:"description,omitempty"`

	// 域名信息列表
	DomainNames []DomainSetInfoDto `json:"domain_names"`

	// 域名组类型，0表示应用域名组，1表示网络域名组
	DomainSetType *int32 `json:"domain_set_type,omitempty"`
}

func (o AddDomainSetInfoDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddDomainSetInfoDto struct{}"
	}

	return strings.Join([]string{"AddDomainSetInfoDto", string(data)}, " ")
}
