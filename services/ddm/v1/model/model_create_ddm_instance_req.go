package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateDdmInstanceReq struct {

	// 实例名称。
	Name *string `json:"name,omitempty"`

	// 可用区。
	AvailableZones *[]string `json:"available_zones,omitempty"`

	// 节点数量。
	NodeNum *int32 `json:"node_num,omitempty"`

	// 引擎版本。
	EngineVersion *string `json:"engine_version,omitempty"`

	// 规格。
	FlavorRef *string `json:"flavor_ref,omitempty"`

	// 虚拟私有云id。
	VpcId *string `json:"vpc_id,omitempty"`

	// 子网id。
	SubnetId *string `json:"subnet_id,omitempty"`

	// 安全组id。
	SecurityGroupId *string `json:"security_group_id,omitempty"`

	// 参数组id。
	ParamGroupId *string `json:"param_group_id,omitempty"`

	// 企业项目id。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 时区。
	TimeZone *string `json:"time_zone,omitempty"`

	// 账号。
	AdminUserName *string `json:"admin_user_name,omitempty"`

	// 密码。
	AdminUserPassword *string `json:"admin_user_password,omitempty"`

	// 付费信息。
	ChargeInfo *interface{} `json:"charge_info,omitempty"`
}

func (o CreateDdmInstanceReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDdmInstanceReq struct{}"
	}

	return strings.Join([]string{"CreateDdmInstanceReq", string(data)}, " ")
}
