package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UnsubscribePostpaidVolumeRequest Request Object
type UnsubscribePostpaidVolumeRequest struct {
	Body *UnsubscribeVolumeRequestBody `json:"body,omitempty"`
}

func (o UnsubscribePostpaidVolumeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UnsubscribePostpaidVolumeRequest struct{}"
	}

	return strings.Join([]string{"UnsubscribePostpaidVolumeRequest", string(data)}, " ")
}
