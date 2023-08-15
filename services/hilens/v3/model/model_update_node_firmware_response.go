package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateNodeFirmwareResponse Response Object
type UpdateNodeFirmwareResponse struct {
	Body *interface{} `json:"body,omitempty"`

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
