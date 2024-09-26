package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AssociateGlobalConnectionBandwidthInstanceRequest Request Object
type AssociateGlobalConnectionBandwidthInstanceRequest struct {

	// 实例ID。
	Id string `json:"id"`

	Body *AssociateGlobalConnectionBandwidthInstanceRequestBody `json:"body,omitempty"`
}

func (o AssociateGlobalConnectionBandwidthInstanceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssociateGlobalConnectionBandwidthInstanceRequest struct{}"
	}

	return strings.Join([]string{"AssociateGlobalConnectionBandwidthInstanceRequest", string(data)}, " ")
}
