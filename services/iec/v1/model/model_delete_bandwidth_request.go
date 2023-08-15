package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteBandwidthRequest Request Object
type DeleteBandwidthRequest struct {

	// 带宽ID。
	BandwidthId string `json:"bandwidth_id"`
}

func (o DeleteBandwidthRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteBandwidthRequest struct{}"
	}

	return strings.Join([]string{"DeleteBandwidthRequest", string(data)}, " ")
}
