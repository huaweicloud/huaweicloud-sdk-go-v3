package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListFirewallDetailRequest Request Object
type ListFirewallDetailRequest struct {

	// 偏移量：指定返回记录的开始位置，必须为数字，取值范围为大于或等于0，默认0
	Offset int32 `json:"offset"`

	// 每页显示个数，范围为1-1024
	Limit int32 `json:"limit"`

	// 服务类型，目前仅支持0互联网防护
	ServiceType int32 `json:"service_type"`

	// 企业项目ID，用户根据组织规划企业项目，对应的ID为企业项目ID，可通过[如何获取企业项目ID](cfw_02_0027.xml)获取，用户未开启企业项目时为0
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 防火墙实例id，创建云防火墙后用于标志防火墙由系统自动生成的标志id，可通过调用[查询防火墙实例接口](ListFirewallDetail.xml)，默认情况下，fw_instance_Id为空时，返回账号下第一个墙的信息；fw_instance_Id非空时，返回与fw_instance_Id对应墙的信息。
	FwInstanceId *string `json:"fw_instance_id,omitempty"`

	// 防火墙名称
	Name *string `json:"name,omitempty"`
}

func (o ListFirewallDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFirewallDetailRequest struct{}"
	}

	return strings.Join([]string{"ListFirewallDetailRequest", string(data)}, " ")
}
