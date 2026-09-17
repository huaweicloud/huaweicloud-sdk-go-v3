package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type InPlaceMigratetoNodesSpec struct {

	// **参数解释**： 腾挪节点列表 **约束限制**： 不涉及
	Nodes []InplaceMigrateNodeItem `json:"nodes"`

	DataDiskCleanUpOption *DataDiskCleanUpOption `json:"dataDiskCleanUpOption,omitempty"`

	ExtendParam *InPlaceMigrateNodeExtendParam `json:"extendParam,omitempty"`
}

func (o InPlaceMigratetoNodesSpec) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InPlaceMigratetoNodesSpec struct{}"
	}

	return strings.Join([]string{"InPlaceMigratetoNodesSpec", string(data)}, " ")
}
