package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListWafBandwidthResponse Response Object
type ListWafBandwidthResponse struct {

	// 曲线
	Curve          *[]Curve `json:"curve,omitempty"`
	HttpStatusCode int      `json:"-"`
}

func (o ListWafBandwidthResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListWafBandwidthResponse struct{}"
	}

	return strings.Join([]string{"ListWafBandwidthResponse", string(data)}, " ")
}
