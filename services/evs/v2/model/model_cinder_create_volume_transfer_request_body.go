package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// This is a auto create Body Object
type CinderCreateVolumeTransferRequestBody struct {
	Transfer *CreateVolumeTransferOption `json:"transfer"`
}

func (o CinderCreateVolumeTransferRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CinderCreateVolumeTransferRequestBody struct{}"
	}

	return strings.Join([]string{"CinderCreateVolumeTransferRequestBody", string(data)}, " ")
}
