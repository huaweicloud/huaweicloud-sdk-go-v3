package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowFlowDetailRequest Request Object
type ShowFlowDetailRequest struct {

	// 流id
	FlowId string `json:"flow_id"`
}

func (o ShowFlowDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowFlowDetailRequest struct{}"
	}

	return strings.Join([]string{"ShowFlowDetailRequest", string(data)}, " ")
}
