package model

import (
	"encoding/json"

	"strings"
)

type MysqlInstanceResponse struct {
	// 实例ID。

	Id string `json:"id"`
	// 实例名称，与请求参数相同。

	Name string `json:"name"`
	// 实例状态。

	Status *string `json:"status,omitempty"`

	Datastore *MysqlDatastore `json:"datastore,omitempty"`
	// 实例类型，与请求参数相同。

	Mode *string `json:"mode,omitempty"`
	// 参数组ID。

	ConfigurationId *string `json:"configuration_id,omitempty"`
	// 数据库端口信息，与请求参数相同。

	Port *string `json:"port,omitempty"`

	BackupStrategy *MysqlBackupStrategy `json:"backup_strategy,omitempty"`
	// 企业项目ID。

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
	// 区域ID，与请求参数相同。

	Region *string `json:"region,omitempty"`
	// 可用区模式，与请求参数相同。

	AvailabilityZoneMode *string `json:"availability_zone_mode,omitempty"`
	// 主可用区ID。

	MasterAvailabilityZone *string `json:"master_availability_zone,omitempty"`
	// 虚拟私有云ID，与请求参数相同。

	VpcId *string `json:"vpc_id,omitempty"`
	// 安全组ID，与请求参数相同。

	SecurityGroupId *string `json:"security_group_id,omitempty"`
	// 子网ID，与请求参数相同。

	SubnetId *string `json:"subnet_id,omitempty"`
	// 规格码，与请求参数相同。

	FlavorRef *string `json:"flavor_ref,omitempty"`

	ChargeInfo *MysqlChargeInfo `json:"charge_info,omitempty"`
}

func (o MysqlInstanceResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "MysqlInstanceResponse struct{}"
	}

	return strings.Join([]string{"MysqlInstanceResponse", string(data)}, " ")
}
