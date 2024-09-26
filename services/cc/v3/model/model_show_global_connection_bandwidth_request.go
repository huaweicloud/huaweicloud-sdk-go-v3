package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowGlobalConnectionBandwidthRequest Request Object
type ShowGlobalConnectionBandwidthRequest struct {

	// 实例ID。
	Id string `json:"id"`
}

func (o ShowGlobalConnectionBandwidthRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowGlobalConnectionBandwidthRequest struct{}"
	}

	return strings.Join([]string{"ShowGlobalConnectionBandwidthRequest", string(data)}, " ")
}
