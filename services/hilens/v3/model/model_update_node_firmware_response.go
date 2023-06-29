package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateNodeFirmwareResponse Response Object
type UpdateNodeFirmwareResponse struct {
	FirmwareName *string `json:"firmware_name,omitempty"`

	FirmwareId     *string `json:"firmware_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateNodeFirmwareResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateNodeFirmwareResponse struct{}"
	}

	return strings.Join([]string{"UpdateNodeFirmwareResponse", string(data)}, " ")
}
