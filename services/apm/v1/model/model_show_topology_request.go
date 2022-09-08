package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowTopologyRequest struct {

	// 调用链tracId
	TraceId string `json:"trace_id"`
}

func (o ShowTopologyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTopologyRequest struct{}"
	}

	return strings.Join([]string{"ShowTopologyRequest", string(data)}, " ")
}
