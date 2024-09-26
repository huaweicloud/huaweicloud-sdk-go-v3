package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListGlobalConnectionBandwidthLineLevelsResponse Response Object
type ListGlobalConnectionBandwidthLineLevelsResponse struct {

	// 请求ID。
	RequestId string `json:"request_id"`

	PageInfo *PageInfo `json:"page_info,omitempty"`

	// 线路分级列表响应体。
	LineLevels     []GlobalConnectionBandwidthLineLevel `json:"line_levels"`
	HttpStatusCode int                                  `json:"-"`
}

func (o ListGlobalConnectionBandwidthLineLevelsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListGlobalConnectionBandwidthLineLevelsResponse struct{}"
	}

	return strings.Join([]string{"ListGlobalConnectionBandwidthLineLevelsResponse", string(data)}, " ")
}
