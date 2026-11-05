package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdatePublicKibanaBandwidthReq struct {
	BandWidth *UpdatePublicKibanaBandwidthReqBandWidth `json:"band_width"`

	IsAutoPay *int32 `json:"is_auto_pay,omitempty"`
}

func (o UpdatePublicKibanaBandwidthReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePublicKibanaBandwidthReq struct{}"
	}

	return strings.Join([]string{"UpdatePublicKibanaBandwidthReq", string(data)}, " ")
}
