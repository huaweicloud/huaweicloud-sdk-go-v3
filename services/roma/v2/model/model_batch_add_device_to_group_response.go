package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchAddDeviceToGroupResponse Response Object
type BatchAddDeviceToGroupResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchAddDeviceToGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchAddDeviceToGroupResponse struct{}"
	}

	return strings.Join([]string{"BatchAddDeviceToGroupResponse", string(data)}, " ")
}
