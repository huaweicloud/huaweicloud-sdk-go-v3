package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteDeviceAuthenticationTemplateResponse Response Object
type DeleteDeviceAuthenticationTemplateResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteDeviceAuthenticationTemplateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDeviceAuthenticationTemplateResponse struct{}"
	}

	return strings.Join([]string{"DeleteDeviceAuthenticationTemplateResponse", string(data)}, " ")
}
