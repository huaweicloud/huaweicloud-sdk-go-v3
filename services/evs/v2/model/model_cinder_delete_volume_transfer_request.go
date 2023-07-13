package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CinderDeleteVolumeTransferRequest Request Object
type CinderDeleteVolumeTransferRequest struct {

	// 云硬盘过户记录ID
	TransferId string `json:"transfer_id"`
}

func (o CinderDeleteVolumeTransferRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CinderDeleteVolumeTransferRequest struct{}"
	}

	return strings.Join([]string{"CinderDeleteVolumeTransferRequest", string(data)}, " ")
}
