package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 集群信息。
type CreateClusterRolesBody struct {

	// 实例规格名称。例如，  - ess.spec-4u32g规格对应的取值范围为40GB～2560GB。 - ess.spec-8u64g规格对应的取值范围为80GB～5120GB。 - ess.spec-16u128g规格对应的取值范围为160GB～10240GB。
	FlavorRef string `json:"flavorRef"`

	Volume *CreateClusterInstanceVolumeBody `json:"volume"`

	// 实例类型(选择实例类型时应至少选择一个ess类型)。例如，  - ess-master对应Master节点。 - ess-client对应Clinet节点。 - ess-cold对应冷数据节点。 - ess对应数据节点。
	Type string `json:"type"`

	// 实例个数。 - ess节点，选择范围：1~32个节点数量。    - 若同时选择ess和ess-master时，可以选择1~200个节点数量。    - 若同时选择ess和ess-client时，可以选择1~32个节点数量。    - 若同时选择ess和ess-cold时，可以选择1~32个节点数量。 - ess-master节点，选择范围：3~10内的奇数个节点数量。 - ess-client节点，选择范围：1~32个节点数量。 - ess-cold节点，选择范围：1~32个节点数量。
	InstanceNum int32 `json:"instanceNum"`
}

func (o CreateClusterRolesBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateClusterRolesBody struct{}"
	}

	return strings.Join([]string{"CreateClusterRolesBody", string(data)}, " ")
}
