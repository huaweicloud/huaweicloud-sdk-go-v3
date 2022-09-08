package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CinderShowVolumeTransferRequest struct {

	// 云硬盘过户记录ID
	TransferId string `json:"transfer_id"`
}

func (o CinderShowVolumeTransferRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CinderShowVolumeTransferRequest struct{}"
	}

	return strings.Join([]string{"CinderShowVolumeTransferRequest", string(data)}, " ")
}
