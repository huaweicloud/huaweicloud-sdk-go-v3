package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 实例对象。
type CreateClusterInstanceBody struct {
	// 实例规格名称。例如乌兰察布三中，  - ess.spec-2u16g规格对应的取值范围为40GB～1280GB。 - ess.spec-4u32g规格对应的取值范围为40GB～2560GB。 - ess.spec-8u64g规格对应的取值范围为80GB～5120GB。 - ess.spec-16u128g规格对应的取值范围为160GB～10240GB。

	FlavorRef string `json:"flavorRef"`

	Volume *CreateClusterInstanceVolumeBody `json:"volume"`

	Nics *CreateClusterInstanceNicsBody `json:"nics"`
	// 可用区。不填时默认创建单AZ。  如果需要创建多AZ，各个AZ之间使用英文逗号分隔，比如az1,az2 ，az不能重复输入，并且要求节点个数大于等于AZ个数。  如果节点个数为AZ个数的倍数，节点将会均匀的分布到各个AZ。如果节点个数不为AZ个数的倍数时，各个AZ分布的节点个数之和的绝对值之差小于等于1。

	AvailabilityZone *string `json:"availability_zone,omitempty"`
}

func (o CreateClusterInstanceBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateClusterInstanceBody struct{}"
	}

	return strings.Join([]string{"CreateClusterInstanceBody", string(data)}, " ")
}
