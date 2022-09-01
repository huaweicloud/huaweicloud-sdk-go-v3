package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type IncidentOrderAuthDetailInfoV2 struct {

	// 授权详情id
	Id *int64 `json:"id,omitempty" xml:"id"`

	// 端口
	Port *int32 `json:"port,omitempty" xml:"port"`

	// 账户
	Account *string `json:"account,omitempty" xml:"account"`

	// 授权详情类型，0控制台 1主机资源
	Type *int32 `json:"type,omitempty" xml:"type"`

	// 实例id
	InstanceId *string `json:"instance_id,omitempty" xml:"instance_id"`

	// 实例名称
	InstanceName *string `json:"instance_name,omitempty" xml:"instance_name"`

	// 区域id
	RegionId *string `json:"region_id,omitempty" xml:"region_id"`
}

func (o IncidentOrderAuthDetailInfoV2) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IncidentOrderAuthDetailInfoV2 struct{}"
	}

	return strings.Join([]string{"IncidentOrderAuthDetailInfoV2", string(data)}, " ")
}
