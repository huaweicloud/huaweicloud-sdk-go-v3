package model

import (
	"encoding/json"

	"strings"
)

type BuildMasterDrRelation struct {
	// 灾备实例的实例ID。
	TargetInstanceId string `json:"target_instance_id"`
	// 灾备实例所在租户的项目 ID。
	TargetProjectId string `json:"target_project_id"`
	// 灾备实例所在的区域 ID。
	TargetRegion string `json:"target_region"`
	// 灾备实例的数据虚拟IP（数据VIP）。
	TargetIp string `json:"target_ip"`
	// 灾备实例的子网地址。
	TargetSubnet string `json:"target_subnet"`
}

func (o BuildMasterDrRelation) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "BuildMasterDrRelation struct{}"
	}

	return strings.Join([]string{"BuildMasterDrRelation", string(data)}, " ")
}
