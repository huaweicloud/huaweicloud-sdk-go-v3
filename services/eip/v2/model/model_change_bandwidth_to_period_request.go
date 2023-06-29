package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangeBandwidthToPeriodRequest Request Object
type ChangeBandwidthToPeriodRequest struct {
	Body *BwChangeToPeriodReq `json:"body,omitempty"`
}

func (o ChangeBandwidthToPeriodRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeBandwidthToPeriodRequest struct{}"
	}

	return strings.Join([]string{"ChangeBandwidthToPeriodRequest", string(data)}, " ")
}
