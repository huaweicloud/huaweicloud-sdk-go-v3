package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PartitionReqBody 集群分区信息
type PartitionReqBody struct {

	// **参数解释**： API类型 **约束限制**： 固定值，不允许修改 **取值范围**： 不涉及 **默认取值**： Partition
	Kind *string `json:"kind,omitempty"`

	// **参数解释**： API版本 **约束限制**： 固定值，不允许修改 **取值范围**： 不涉及 **默认取值**： v3
	ApiVersion *string `json:"apiVersion,omitempty"`

	Metadata *PartitionReqBodyMetadata `json:"metadata,omitempty"`

	Spec *PartitionReqBodySpec `json:"spec,omitempty"`
}

func (o PartitionReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PartitionReqBody struct{}"
	}

	return strings.Join([]string{"PartitionReqBody", string(data)}, " ")
}
