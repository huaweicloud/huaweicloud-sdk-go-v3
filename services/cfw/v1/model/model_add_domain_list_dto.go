package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AddDomainListDto struct {

	// 防火墙实例id，创建云防火墙后用于标志防火墙由系统自动生成的标志id，可通过调用[查询防火墙实例接口](ListFirewallDetail.xml)，默认情况下，fw_instance_Id为空时，返回账号下第一个墙的信息；fw_instance_Id非空时，返回与fw_instance_Id对应墙的信息。
	FwInstanceId string `json:"fw_instance_id"`

	// 互联网边界防护对象id，可通过调用[查询防火墙实例接口](ListFirewallDetail.xml)，type为0的为互联网边界防护对象id。
	ObjectId string `json:"object_id"`

	// 域名列表
	DomainNames []DomainSetInfoDto `json:"domain_names"`
}

func (o AddDomainListDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddDomainListDto struct{}"
	}

	return strings.Join([]string{"AddDomainListDto", string(data)}, " ")
}
