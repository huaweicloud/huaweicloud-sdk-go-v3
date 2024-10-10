package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDDoSFlowResponse Response Object
type ListDDoSFlowResponse struct {

	// 当请求type=bps时必返回
	FlowBps *[]FlowBps `json:"flow_bps,omitempty"`

	// 当请求type=pps时必返回
	FlowPps        *[]FlowPps `json:"flow_pps,omitempty"`
	HttpStatusCode int        `json:"-"`
}

func (o ListDDoSFlowResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDDoSFlowResponse struct{}"
	}

	return strings.Join([]string{"ListDDoSFlowResponse", string(data)}, " ")
}
