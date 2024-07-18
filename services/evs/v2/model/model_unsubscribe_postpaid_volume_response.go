package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UnsubscribePostpaidVolumeResponse Response Object
type UnsubscribePostpaidVolumeResponse struct {
	Body           *[]UnsubscribeVolumeResponseBody `json:"body,omitempty"`
	HttpStatusCode int                              `json:"-"`
}

func (o UnsubscribePostpaidVolumeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UnsubscribePostpaidVolumeResponse struct{}"
	}

	return strings.Join([]string{"UnsubscribePostpaidVolumeResponse", string(data)}, " ")
}
