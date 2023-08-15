package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddPublicipsIntoSharedBandwidthRequest Request Object
type AddPublicipsIntoSharedBandwidthRequest struct {

	// 带宽唯一标识
	BandwidthId string `json:"bandwidth_id"`

	Body *AddPublicipsIntoSharedBandwidthRequestBody `json:"body,omitempty"`
}

func (o AddPublicipsIntoSharedBandwidthRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddPublicipsIntoSharedBandwidthRequest struct{}"
	}

	return strings.Join([]string{"AddPublicipsIntoSharedBandwidthRequest", string(data)}, " ")
}
