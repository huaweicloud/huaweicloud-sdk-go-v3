package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSlowLogDetailSampleResponse Response Object
type ShowSlowLogDetailSampleResponse struct {
	Sample         *SlowLogDetail `json:"sample,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ShowSlowLogDetailSampleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSlowLogDetailSampleResponse struct{}"
	}

	return strings.Join([]string{"ShowSlowLogDetailSampleResponse", string(data)}, " ")
}
