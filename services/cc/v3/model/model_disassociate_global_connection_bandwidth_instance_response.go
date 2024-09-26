package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DisassociateGlobalConnectionBandwidthInstanceResponse Response Object
type DisassociateGlobalConnectionBandwidthInstanceResponse struct {

	// 请求ID。
	RequestId string `json:"request_id"`

	// 全域互联带宽解绑实例响应详情。
	Gcbandwidths   []DisassociateGlobalConnectionBandwidthInstanceResponseInfo `json:"gcbandwidths"`
	HttpStatusCode int                                                         `json:"-"`
}

func (o DisassociateGlobalConnectionBandwidthInstanceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DisassociateGlobalConnectionBandwidthInstanceResponse struct{}"
	}

	return strings.Join([]string{"DisassociateGlobalConnectionBandwidthInstanceResponse", string(data)}, " ")
}
