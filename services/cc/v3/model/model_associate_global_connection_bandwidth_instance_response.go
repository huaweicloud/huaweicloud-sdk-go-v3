package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AssociateGlobalConnectionBandwidthInstanceResponse Response Object
type AssociateGlobalConnectionBandwidthInstanceResponse struct {

	// 请求ID。
	RequestId string `json:"request_id"`

	// 全域互联带宽绑定实例响应详情。
	Gcbandwidths   []AssociateGlobalConnectionBandwidthInstanceResponseInfo `json:"gcbandwidths"`
	HttpStatusCode int                                                      `json:"-"`
}

func (o AssociateGlobalConnectionBandwidthInstanceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssociateGlobalConnectionBandwidthInstanceResponse struct{}"
	}

	return strings.Join([]string{"AssociateGlobalConnectionBandwidthInstanceResponse", string(data)}, " ")
}
