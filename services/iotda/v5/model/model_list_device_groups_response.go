package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDeviceGroupsResponse Response Object
type ListDeviceGroupsResponse struct {

	// 设备组信息列表。
	DeviceGroups *[]DeviceGroupResponseSummary `json:"device_groups,omitempty"`

	Page           *Page `json:"page,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o ListDeviceGroupsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDeviceGroupsResponse struct{}"
	}

	return strings.Join([]string{"ListDeviceGroupsResponse", string(data)}, " ")
}
