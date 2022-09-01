package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ShowCeshierarchyRespQueues struct {

	// topic名称。
	Name *string `json:"name,omitempty" xml:"name"`

	// 分区列表。
	Partitions *[]ShowCeshierarchyRespPartitions `json:"partitions,omitempty" xml:"partitions"`
}

func (o ShowCeshierarchyRespQueues) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCeshierarchyRespQueues struct{}"
	}

	return strings.Join([]string{"ShowCeshierarchyRespQueues", string(data)}, " ")
}
