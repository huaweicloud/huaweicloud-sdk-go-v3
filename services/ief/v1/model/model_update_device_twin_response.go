package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateDeviceTwinResponse Response Object
type UpdateDeviceTwinResponse struct {

	// 孪生属性配置
	PropertyVisitors map[string]ValueInPropertyVisitors `json:"property_visitors,omitempty"`

	// 终端设备动态属性
	Twin           map[string]ValueInTwinResponse `json:"twin,omitempty"`
	HttpStatusCode int                            `json:"-"`
}

func (o UpdateDeviceTwinResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDeviceTwinResponse struct{}"
	}

	return strings.Join([]string{"UpdateDeviceTwinResponse", string(data)}, " ")
}
