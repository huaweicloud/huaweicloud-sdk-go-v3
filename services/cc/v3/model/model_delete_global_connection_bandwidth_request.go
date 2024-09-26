package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteGlobalConnectionBandwidthRequest Request Object
type DeleteGlobalConnectionBandwidthRequest struct {

	// 实例ID。
	Id string `json:"id"`
}

func (o DeleteGlobalConnectionBandwidthRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteGlobalConnectionBandwidthRequest struct{}"
	}

	return strings.Join([]string{"DeleteGlobalConnectionBandwidthRequest", string(data)}, " ")
}
