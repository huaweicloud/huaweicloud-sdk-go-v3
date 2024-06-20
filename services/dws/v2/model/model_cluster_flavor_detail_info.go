package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ClusterFlavorDetailInfo struct {

	// 规格ID
	Id string `json:"id"`

	// 规格编码
	SpecName string `json:"spec_name"`

	// 当前节点数量
	CurrentNode int32 `json:"current_node"`

	// 最小节点阈值
	MinNode int32 `json:"min_node"`

	// 最大节点阈值
	MaxNode int32 `json:"max_node"`

	// 规格类型
	Classify string `json:"classify"`

	// 数据仓库版本
	DatastoreVersion string `json:"datastore_version"`

	// 扩展信息
	Attribute []FlavorAttributeInfo `json:"attribute"`

	VolumeNode *FlavorVolumeNodeInfo `json:"volume_node"`
}

func (o ClusterFlavorDetailInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ClusterFlavorDetailInfo struct{}"
	}

	return strings.Join([]string{"ClusterFlavorDetailInfo", string(data)}, " ")
}
