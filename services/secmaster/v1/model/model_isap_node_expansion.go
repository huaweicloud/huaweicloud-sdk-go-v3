package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IsapNodeExpansion 节点扩展信息
type IsapNodeExpansion struct {

	// 自定义标签
	CustomLabel *string `json:"custom_label,omitempty"`

	// 数据中心
	DataCenter *string `json:"data_center,omitempty"`

	// 描述信息
	Description *string `json:"description,omitempty"`

	// 维护人
	Maintainer *string `json:"maintainer,omitempty"`

	// 网络平面
	NetworkPlane *string `json:"network_plane,omitempty"`

	// UUID
	NodeId *string `json:"node_id,omitempty"`
}

func (o IsapNodeExpansion) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IsapNodeExpansion struct{}"
	}

	return strings.Join([]string{"IsapNodeExpansion", string(data)}, " ")
}
