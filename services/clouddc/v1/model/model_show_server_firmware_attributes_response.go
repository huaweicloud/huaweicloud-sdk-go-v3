package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowServerFirmwareAttributesResponse Response Object
type ShowServerFirmwareAttributesResponse struct {

	// 固件详细信息列表
	Body           *[]FirmwareDetails `json:"body,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o ShowServerFirmwareAttributesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowServerFirmwareAttributesResponse struct{}"
	}

	return strings.Join([]string{"ShowServerFirmwareAttributesResponse", string(data)}, " ")
}
