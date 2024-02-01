package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AssociateGlobalConnectionBandwidthInstanceRequestBody 全域互联带宽绑定实例详细信息。
type AssociateGlobalConnectionBandwidthInstanceRequestBody struct {

	// 全域互联带宽绑定实例详细信息。
	Gcbandwidths []AssociateGlobalConnectionBandwidthInstanceRequestInfo `json:"gcbandwidths"`
}

func (o AssociateGlobalConnectionBandwidthInstanceRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssociateGlobalConnectionBandwidthInstanceRequestBody struct{}"
	}

	return strings.Join([]string{"AssociateGlobalConnectionBandwidthInstanceRequestBody", string(data)}, " ")
}
