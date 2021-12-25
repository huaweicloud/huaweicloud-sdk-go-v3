package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
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
