package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowServerBlockDeviceResponse Response Object
type ShowServerBlockDeviceResponse struct {
	VolumeAttachment *ServerBlockDevice `json:"volumeAttachment,omitempty"`
	HttpStatusCode   int                `json:"-"`
}

func (o ShowServerBlockDeviceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowServerBlockDeviceResponse struct{}"
	}

	return strings.Join([]string{"ShowServerBlockDeviceResponse", string(data)}, " ")
}
