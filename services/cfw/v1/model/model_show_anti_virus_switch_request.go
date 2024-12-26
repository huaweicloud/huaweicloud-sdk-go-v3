package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAntiVirusSwitchRequest Request Object
type ShowAntiVirusSwitchRequest struct {

	// 防护对象ID，是创建云防火墙后用于区分互联网边界防护和VPC边界防护的标志id，可通过调用[查询防火墙实例接口](ListFirewallDetail.xml)获得，通过返回值中的data.records.protect_objects.object_id（.表示各对象之间层级的区分）获得，注意type为0的为互联网边界防护对象id，type为1的为VPC边界防护对象id。此处仅取type为1的防护对象id，可通过data.records.protect_objects.type（.表示各对象之间层级的区分）获得。
	ObjectId string `json:"object_id"`

	// 企业项目ID，用户根据组织规划企业项目，对应的ID为企业项目ID，可通过[如何获取企业项目ID](cfw_02_0027.xml)获取，用户未开启企业项目时为0
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
}

func (o ShowAntiVirusSwitchRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAntiVirusSwitchRequest struct{}"
	}

	return strings.Join([]string{"ShowAntiVirusSwitchRequest", string(data)}, " ")
}
