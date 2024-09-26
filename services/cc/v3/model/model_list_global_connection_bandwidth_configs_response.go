package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListGlobalConnectionBandwidthConfigsResponse Response Object
type ListGlobalConnectionBandwidthConfigsResponse struct {

	// 请求ID。
	RequestId string `json:"request_id"`

	Configs        *ListGlobalConnectionBandwidthConfigs `json:"configs,omitempty"`
	HttpStatusCode int                                   `json:"-"`
}

func (o ListGlobalConnectionBandwidthConfigsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListGlobalConnectionBandwidthConfigsResponse struct{}"
	}

	return strings.Join([]string{"ListGlobalConnectionBandwidthConfigsResponse", string(data)}, " ")
}
