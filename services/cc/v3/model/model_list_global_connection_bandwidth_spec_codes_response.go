package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListGlobalConnectionBandwidthSpecCodesResponse Response Object
type ListGlobalConnectionBandwidthSpecCodesResponse struct {

	// 请求ID。
	RequestId string `json:"request_id"`

	PageInfo *PageInfo `json:"page_info,omitempty"`

	// 线路规格列表响应体。
	SpecCodes      []GlobalConnectionBandwidthSpecCode `json:"spec_codes"`
	HttpStatusCode int                                 `json:"-"`
}

func (o ListGlobalConnectionBandwidthSpecCodesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListGlobalConnectionBandwidthSpecCodesResponse struct{}"
	}

	return strings.Join([]string{"ListGlobalConnectionBandwidthSpecCodesResponse", string(data)}, " ")
}
