package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ClusterDetailInstanceFlavor 节点的虚拟机规格，请参见flavor参数说明（查询集群列表时返回值为null）。
type ClusterDetailInstanceFlavor struct {

	// 节点虚拟机的规格ID。
	Id *string `json:"id,omitempty"`

	// 链接信息
	Links *[]ClusterLinks `json:"links,omitempty"`
}

func (o ClusterDetailInstanceFlavor) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ClusterDetailInstanceFlavor struct{}"
	}

	return strings.Join([]string{"ClusterDetailInstanceFlavor", string(data)}, " ")
}
