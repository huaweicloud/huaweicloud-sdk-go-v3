package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteGlobalConnectionBandwidthResponse Response Object
type DeleteGlobalConnectionBandwidthResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteGlobalConnectionBandwidthResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteGlobalConnectionBandwidthResponse struct{}"
	}

	return strings.Join([]string{"DeleteGlobalConnectionBandwidthResponse", string(data)}, " ")
}
