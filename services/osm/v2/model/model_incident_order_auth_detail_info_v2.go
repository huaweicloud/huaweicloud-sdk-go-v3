package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type IncidentOrderAuthDetailInfoV2 struct {
	// 授权详情id

	Id *int64 `json:"id,omitempty"`
	// 端口

	Port *int32 `json:"port,omitempty"`
	// 账户

	Account *string `json:"account,omitempty"`
	// 授权详情类型，0控制台 1主机资源

	Type *int32 `json:"type,omitempty"`
	// 实例id

	InstanceId *string `json:"instance_id,omitempty"`
	// 实例名称

	InstanceName *string `json:"instance_name,omitempty"`
	// 区域id

	RegionId *string `json:"region_id,omitempty"`
}

func (o IncidentOrderAuthDetailInfoV2) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IncidentOrderAuthDetailInfoV2 struct{}"
	}

	return strings.Join([]string{"IncidentOrderAuthDetailInfoV2", string(data)}, " ")
}
