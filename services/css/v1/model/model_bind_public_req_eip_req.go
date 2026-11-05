package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BindPublicReqEipReq 弹性IP信息。
type BindPublicReqEipReq struct {
	BandWidth *BindPublicReqEipBandWidth `json:"band_width"`

	IsAutoPay *int32 `json:"is_auto_pay,omitempty"`
}

func (o BindPublicReqEipReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BindPublicReqEipReq struct{}"
	}

	return strings.Join([]string{"BindPublicReqEipReq", string(data)}, " ")
}
