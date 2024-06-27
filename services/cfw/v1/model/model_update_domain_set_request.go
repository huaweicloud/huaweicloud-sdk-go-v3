package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateDomainSetRequest Request Object
type UpdateDomainSetRequest struct {

	// 域名组id
	SetId string `json:"set_id"`

	// 企业项目id，用户支持企业项目后，由企业项目生成的id。
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
