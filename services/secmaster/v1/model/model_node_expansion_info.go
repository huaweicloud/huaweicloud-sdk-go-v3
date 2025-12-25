package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type NodeExpansionInfo struct {

	// 自定义标签
	CustomLabel *string `json:"custom_label,omitempty"`

	// 数据中心
	DataCenter *string `json:"data_center,omitempty"`

	// 描述
	Description *string `json:"description,omitempty"`

	// 维护人
	Maintainer *string `json:"maintainer,omitempty"`

	// 网络平面
	NetworkPlane *string `json:"network_plane,omitempty"`
}

func (o NodeExpansionInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NodeExpansionInfo struct{}"
	}

	return strings.Join([]string{"NodeExpansionInfo", string(data)}, " ")
}
