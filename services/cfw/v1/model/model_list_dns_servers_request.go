package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDnsServersRequest Request Object
type ListDnsServersRequest struct {

	// 每页显示个数，范围为1-1024
	Limit *int32 `json:"limit,omitempty"`

	// 偏移量：指定返回记录的开始位置，必须为数字，取值范围为大于或等于0，默认0
	Offset *int32 `json:"offset,omitempty"`

	// 防火墙实例id，创建云防火墙后用于标志防火墙由系统自动生成的标志id，可通过调用[查询防火墙实例接口](ListFirewallDetail.xml)接口获得，默认情况下，fw_instance_Id为空时，返回账号下第一个墙的信息；fw_instance_Id非空时，返回与fw_instance_Id对应墙的信息。若object_Id非空，默认返回object_Id对应墙的信息；填写时object_Id需要属于fw_instance_Id对应的墙。
	FwInstanceId string `json:"fw_instance_id"`

	// 企业项目id，用户支持企业项目后，由企业项目生成的id。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
}

func (o ListDnsServersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDnsServersRequest struct{}"
	}

	return strings.Join([]string{"ListDnsServersRequest", string(data)}, " ")
}
