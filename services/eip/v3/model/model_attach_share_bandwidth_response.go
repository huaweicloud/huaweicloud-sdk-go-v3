package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AttachShareBandwidthResponse Response Object
type AttachShareBandwidthResponse struct {
	Publicip *PublicipResp `json:"publicip,omitempty"`

	// 本次请求编号
	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o AttachShareBandwidthResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AttachShareBandwidthResponse struct{}"
	}

	return strings.Join([]string{"AttachShareBandwidthResponse", string(data)}, " ")
}
