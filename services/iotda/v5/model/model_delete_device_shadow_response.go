package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteDeviceShadowResponse Response Object
type DeleteDeviceShadowResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteDeviceShadowResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDeviceShadowResponse struct{}"
	}

	return strings.Join([]string{"DeleteDeviceShadowResponse", string(data)}, " ")
}
