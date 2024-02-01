package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DisassociateGlobalConnectionBandwidthInstanceRequestBody 全域互联带宽解绑实例详细信息。
type DisassociateGlobalConnectionBandwidthInstanceRequestBody struct {

	// 全域互联带宽解绑实例详细信息。
	Gcbandwidths []DisassociateGlobalConnectionBandwidthInstanceRequestInfo `json:"gcbandwidths"`
}

func (o DisassociateGlobalConnectionBandwidthInstanceRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DisassociateGlobalConnectionBandwidthInstanceRequestBody struct{}"
	}

	return strings.Join([]string{"DisassociateGlobalConnectionBandwidthInstanceRequestBody", string(data)}, " ")
}
