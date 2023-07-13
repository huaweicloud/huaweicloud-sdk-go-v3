package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CinderCreateVolumeTransferRequest Request Object
type CinderCreateVolumeTransferRequest struct {
	Body *CinderCreateVolumeTransferRequestBody `json:"body,omitempty"`
}

func (o CinderCreateVolumeTransferRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CinderCreateVolumeTransferRequest struct{}"
	}

	return strings.Join([]string{"CinderCreateVolumeTransferRequest", string(data)}, " ")
}
