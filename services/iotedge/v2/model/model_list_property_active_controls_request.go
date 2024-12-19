package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPropertyActiveControlsRequest Request Object
type ListPropertyActiveControlsRequest struct {

	// 边缘节点ID
	EdgeNodeId string `json:"edge_node_id"`

	// 设备ID
	DeviceId string `json:"device_id"`

	// 设备服务id。可选，在属性平铺场景不需要填，如果不填则表示service_id为空字符串
	ServiceId *string `json:"service_id,omitempty"`

	// 设备属性。必选
	Property string `json:"property"`
}

func (o ListPropertyActiveControlsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPropertyActiveControlsRequest struct{}"
	}

	return strings.Join([]string{"ListPropertyActiveControlsRequest", string(data)}, " ")
}
