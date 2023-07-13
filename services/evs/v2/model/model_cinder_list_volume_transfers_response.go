package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CinderListVolumeTransfersResponse Response Object
type CinderListVolumeTransfersResponse struct {

	// 云硬盘过户记录列表概要，请参见•[transfers参数说明](https://support.huaweicloud.com/api-evs/evs_04_2110.html#evs_04_2110__li6113282511345)。
	Transfers      *[]VolumeTransferSummary `json:"transfers,omitempty"`
	HttpStatusCode int                      `json:"-"`
}

func (o CinderListVolumeTransfersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CinderListVolumeTransfersResponse struct{}"
	}

	return strings.Join([]string{"CinderListVolumeTransfersResponse", string(data)}, " ")
}
