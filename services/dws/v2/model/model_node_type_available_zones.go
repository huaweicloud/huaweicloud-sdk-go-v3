package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 节点类型支持的可用区及状态信息。
type NodeTypeAvailableZones struct {

	// 可用区ID。
	Code string `json:"code"`

	// 节点类型可用状态。 - normal：可用 - sellout：售罄 - abandon：不可用
	Status string `json:"status"`
}

func (o NodeTypeAvailableZones) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NodeTypeAvailableZones struct{}"
	}

	return strings.Join([]string{"NodeTypeAvailableZones", string(data)}, " ")
}
