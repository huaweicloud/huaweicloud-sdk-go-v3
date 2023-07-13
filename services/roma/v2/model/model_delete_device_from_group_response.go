package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteDeviceFromGroupResponse Response Object
type DeleteDeviceFromGroupResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteDeviceFromGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDeviceFromGroupResponse struct{}"
	}

	return strings.Join([]string{"DeleteDeviceFromGroupResponse", string(data)}, " ")
}
