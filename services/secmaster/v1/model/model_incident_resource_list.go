package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type IncidentResourceList struct {

	// 云服务资产id
	Id *string `json:"id,omitempty"`

	// 资产名称
	Name *string `json:"name,omitempty"`

	// 资产类型(cloudservers;servers;vm;pm;instances;publicips;ip;website;vpcs;securityGroups;device)
	Type *string `json:"type,omitempty"`

	// 云服务名称；引用云RMS provider字段
	Provider *string `json:"provider,omitempty"`

	// 区域；按照云regionId填写
	RegionId *string `json:"region_id,omitempty"`

	// 资产所属账号ID，UUID格式
	DomainId *string `json:"domain_id,omitempty"`

	// 资产所属项目ID，UUID格式
	ProjectId *string `json:"project_id,omitempty"`

	// 企业项目id
	EpId *string `json:"ep_id,omitempty"`

	// 企业项目名称
	EpName *string `json:"ep_name,omitempty"`

	// 资产标签 1、最多50个key/values对 2、values：最大255字符，取值范围：字母数字,空格,+, -, =, ., _, :, /,@
	Tags *string `json:"tags,omitempty"`
}

func (o IncidentResourceList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IncidentResourceList struct{}"
	}

	return strings.Join([]string{"IncidentResourceList", string(data)}, " ")
}
