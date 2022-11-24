package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CinderAcceptVolumeTransferResponse struct {
	Transfer       *VolumeTransferSummary `json:"transfer,omitempty"`
	HttpStatusCode int                    `json:"-"`
}

func (o CinderAcceptVolumeTransferResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CinderAcceptVolumeTransferResponse struct{}"
	}

	return strings.Join([]string{"CinderAcceptVolumeTransferResponse", string(data)}, " ")
}
