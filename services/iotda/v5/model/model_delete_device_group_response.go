package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DeleteDeviceGroupResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteDeviceGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDeviceGroupResponse struct{}"
	}

	return strings.Join([]string{"DeleteDeviceGroupResponse", string(data)}, " ")
}
