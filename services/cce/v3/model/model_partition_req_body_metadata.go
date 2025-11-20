package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PartitionReqBodyMetadata **参数解释**： 分区的元数据信息 **约束限制**： 不涉及
type PartitionReqBodyMetadata struct {

	// **参数解释**： 分区名称 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	Name *string `json:"name,omitempty"`
}

func (o PartitionReqBodyMetadata) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PartitionReqBodyMetadata struct{}"
	}

	return strings.Join([]string{"PartitionReqBodyMetadata", string(data)}, " ")
}
