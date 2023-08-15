package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteDeviceGroupResponse Response Object
type DeleteDeviceGroupResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteDeviceGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDeviceGroupResponse struct{}"
	}

	return strings.Join([]string{"DeleteDeviceGroupResponse", string(data)}, " ")
}
