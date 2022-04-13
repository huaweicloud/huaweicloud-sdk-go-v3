package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DeleteDeviceTemplateResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteDeviceTemplateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDeviceTemplateResponse struct{}"
	}

	return strings.Join([]string{"DeleteDeviceTemplateResponse", string(data)}, " ")
}
