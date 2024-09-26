package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DisassociateGlobalConnectionBandwidthInstanceRequest Request Object
type DisassociateGlobalConnectionBandwidthInstanceRequest struct {

	// 实例ID。
	Id string `json:"id"`

	Body *DisassociateGlobalConnectionBandwidthInstanceRequestBody `json:"body,omitempty"`
}

func (o DisassociateGlobalConnectionBandwidthInstanceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DisassociateGlobalConnectionBandwidthInstanceRequest struct{}"
	}

	return strings.Join([]string{"DisassociateGlobalConnectionBandwidthInstanceRequest", string(data)}, " ")
}
