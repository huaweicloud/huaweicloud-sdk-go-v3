package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateGlobalConnectionBandwidthRequest Request Object
type UpdateGlobalConnectionBandwidthRequest struct {

	// 实例ID。
	Id string `json:"id"`

	Body *UpdateGlobalConnectionBandwidthRequestBody `json:"body,omitempty"`
}

func (o UpdateGlobalConnectionBandwidthRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateGlobalConnectionBandwidthRequest struct{}"
	}

	return strings.Join([]string{"UpdateGlobalConnectionBandwidthRequest", string(data)}, " ")
}
