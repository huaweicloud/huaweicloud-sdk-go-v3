package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateGlobalConnectionBandwidthRequest Request Object
type CreateGlobalConnectionBandwidthRequest struct {
	Body *CreateGlobalConnectionBandwidthRequestBody `json:"body,omitempty"`
}

func (o CreateGlobalConnectionBandwidthRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateGlobalConnectionBandwidthRequest struct{}"
	}

	return strings.Join([]string{"CreateGlobalConnectionBandwidthRequest", string(data)}, " ")
}
