package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResourceDetail 资产详情
type ResourceDetail struct {

	// 资产id
	Id string `json:"id"`

	// 资产名称
	Name string `json:"name"`

	// 资产来源，云服务名称(云上)，线下机房（IDC）
	Provider string `json:"provider"`

	// 资产类型, 比如ECS/VPC/EVS/IP/URL等
	Type string `json:"type"`

	// 资产详情校验码。
	Checksum *string `json:"checksum,omitempty"`

	// 资产创建时间。
	Created *string `json:"created,omitempty"`

	// 资产操作状态。
	ProvisioningState *string `json:"provisioning_state,omitempty"`

	Environment *ResourceEnvironment `json:"environment"`

	Department *Department `json:"department,omitempty"`

	GovernanceUser *GovernanceUser `json:"governance_user,omitempty"`

	// 0: 测试 1： 一般   2： 关键资产
	Level *int32 `json:"level,omitempty"`

	Properties *Properties `json:"properties"`
}

func (o ResourceDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceDetail struct{}"
	}

	return strings.Join([]string{"ResourceDetail", string(data)}, " ")
}
