package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDevicesInGroupResponse Response Object
type ShowDevicesInGroupResponse struct {

	// 总数
	Total *int32 `json:"total,omitempty"`

	// 本次返回数量
	Size *int32 `json:"size,omitempty"`

	// 组内设备清单
	Items          *[]DevicesInGroup `json:"items,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o ShowDevicesInGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDevicesInGroupResponse struct{}"
	}

	return strings.Join([]string{"ShowDevicesInGroupResponse", string(data)}, " ")
}
