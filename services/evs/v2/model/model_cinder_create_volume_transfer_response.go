package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CinderCreateVolumeTransferResponse Response Object
type CinderCreateVolumeTransferResponse struct {
	Transfer       *CreateVolumeTransferDetail `json:"transfer,omitempty"`
	HttpStatusCode int                         `json:"-"`
}

func (o CinderCreateVolumeTransferResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CinderCreateVolumeTransferResponse struct{}"
	}

	return strings.Join([]string{"CinderCreateVolumeTransferResponse", string(data)}, " ")
}
