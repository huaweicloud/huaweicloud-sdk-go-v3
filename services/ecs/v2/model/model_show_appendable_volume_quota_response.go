package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAppendableVolumeQuotaResponse Response Object
type ShowAppendableVolumeQuotaResponse struct {
	FreeScsi *int32 `json:"free_scsi,omitempty"`

	Count *int32 `json:"count,omitempty"`

	FreeBlk        *int32 `json:"free_blk,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowAppendableVolumeQuotaResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAppendableVolumeQuotaResponse struct{}"
	}

	return strings.Join([]string{"ShowAppendableVolumeQuotaResponse", string(data)}, " ")
}
