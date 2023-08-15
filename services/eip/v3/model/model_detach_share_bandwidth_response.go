package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DetachShareBandwidthResponse Response Object
type DetachShareBandwidthResponse struct {
	Publicip *PublicipResp `json:"publicip,omitempty"`

	// 本次请求编号
	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DetachShareBandwidthResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DetachShareBandwidthResponse struct{}"
	}

	return strings.Join([]string{"DetachShareBandwidthResponse", string(data)}, " ")
}
