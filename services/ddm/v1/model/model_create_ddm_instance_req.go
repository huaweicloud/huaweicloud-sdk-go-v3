package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateDdmInstanceReq struct {

	// 实例名称。
	Name string `json:"name"`

	// 可用区。
	AvailableZones []string `json:"available_zones"`

	// 节点数量。
	NodeNum int32 `json:"node_num"`

	// 引擎版本。
	EngineVersion string `json:"engine_version"`

	// 规格。
	FlavorRef string `json:"flavor_ref"`

	// 虚拟私有云id。
	VpcId string `json:"vpc_id"`

	// 子网id。
	SubnetId string `json:"subnet_id"`

	// 安全组id。
	SecurityGroupId string `json:"security_group_id"`

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

	ChargeInfo *ChargeInfo `json:"charge_info,omitempty"`
}

func (o CreateDdmInstanceReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDdmInstanceReq struct{}"
	}

	return strings.Join([]string{"CreateDdmInstanceReq", string(data)}, " ")
}
