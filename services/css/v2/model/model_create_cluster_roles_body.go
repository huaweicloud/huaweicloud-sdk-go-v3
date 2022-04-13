package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 集群信息
type CreateClusterRolesBody struct {
	// 实例规格名称。例如，  - ess.spec-2u16g规格对应的取值范围为40GB～1280GB。 - ess.spec-4u32g规格对应的取值范围为40GB～2560GB。 - ess.spec-8u64g规格对应的取值范围为80GB～5120GB。 - ess.spec-16u128g规格对应的取值范围为160GB～10240GB。

	FlavorRef string `json:"flavorRef"`

	Volume *CreateClusterInstanceVolumeBody `json:"volume"`
	// 实例类型。例如，  - ess-master对应Master节点。 - ess-client对应Clinet节点。 - ess-cold对应冷数据节点。 - ess对应数据节点。

	Type string `json:"type"`
	// 实例个数。

	InstanceNum int32 `json:"instanceNum"`
}

func (o CreateClusterRolesBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateClusterRolesBody struct{}"
	}

	return strings.Join([]string{"CreateClusterRolesBody", string(data)}, " ")
}
