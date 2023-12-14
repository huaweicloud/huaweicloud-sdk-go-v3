package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateStreamRequest Request Object
type UpdateStreamRequest struct {

	// 需要变更分区数量的通道名称。
	StreamName string `json:"stream_name"`

	Body *UpdateStreamReq `json:"body,omitempty"`
}

func (o UpdateStreamRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateStreamRequest struct{}"
	}

	return strings.Join([]string{"UpdateStreamRequest", string(data)}, " ")
}
