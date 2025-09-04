package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAppendableVolumeQuotaRequest Request Object
type ShowAppendableVolumeQuotaRequest struct {

	// 云服务器ID。
	ServerId string `json:"server_id"`
}

func (o ShowAppendableVolumeQuotaRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAppendableVolumeQuotaRequest struct{}"
	}

	return strings.Join([]string{"ShowAppendableVolumeQuotaRequest", string(data)}, " ")
}
