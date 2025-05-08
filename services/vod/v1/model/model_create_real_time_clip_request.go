package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateRealTimeClipRequest Request Object
type CreateRealTimeClipRequest struct {
	Body *CreateRealTimeClipRequestBody `json:"body,omitempty"`
}

func (o CreateRealTimeClipRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateRealTimeClipRequest struct{}"
	}

	return strings.Join([]string{"CreateRealTimeClipRequest", string(data)}, " ")
}
