package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 节点的虚拟机规格，请参见flavor参数说明。
type ClusterDetailInstanceFlavor struct {
	// 节点虚拟机的规格ID。

	Id *string `json:"id,omitempty"`
}

func (o ClusterDetailInstanceFlavor) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ClusterDetailInstanceFlavor struct{}"
	}

	return strings.Join([]string{"ClusterDetailInstanceFlavor", string(data)}, " ")
}
