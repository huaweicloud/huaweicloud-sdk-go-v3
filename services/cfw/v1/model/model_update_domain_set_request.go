package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateDomainSetRequest Request Object
type UpdateDomainSetRequest struct {

	// 域名组id，可通过[查询域名组列表接口](ListDomainSets.xml)查询获得，通过返回值中的data.records.set_id（.表示各对象之间层级的区分）获得。
	SetId string `json:"set_id"`

	// 企业项目ID，用户根据组织规划企业项目，对应的ID为企业项目ID，可通过[如何获取企业项目ID](cfw_02_0027.xml)获取，用户未开启企业项目时为0
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 防火墙实例id，创建云防火墙后用于标志防火墙由系统自动生成的标志id，可通过调用[查询防火墙实例接口](ListFirewallDetail.xml)。
	FwInstanceId string `json:"fw_instance_id"`

	Body *UpdateDomainSetInfoDto `json:"body,omitempty"`
}

func (o UpdateDomainSetRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDomainSetRequest struct{}"
	}

	return strings.Join([]string{"UpdateDomainSetRequest", string(data)}, " ")
}
