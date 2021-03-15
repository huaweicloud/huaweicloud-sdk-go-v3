package model

import (
	"encoding/json"

	"strings"
)

type BuildSlaveDrRelation struct {
	// 主实例的实例 ID。
	TargetInstanceId string `json:"target_instance_id"`
	// 主实例所在租户的项目 ID。
	TargetProjectId string `json:"target_project_id"`
	// 主实例所在的区域 ID。
	TargetRegion string `json:"target_region"`
	// 主实例的数据虚拟IP（数据VIP）。
	TargetIp string `json:"target_ip"`
}

func (o BuildSlaveDrRelation) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "BuildSlaveDrRelation struct{}"
	}

	return strings.Join([]string{"BuildSlaveDrRelation", string(data)}, " ")
}
