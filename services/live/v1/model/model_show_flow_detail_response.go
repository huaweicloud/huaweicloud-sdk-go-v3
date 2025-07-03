package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowFlowDetailResponse Response Object
type ShowFlowDetailResponse struct {

	// 流名称
	Name *string `json:"name,omitempty"`

	// 区域
	Region *string `json:"region,omitempty"`

	// 入流信息
	Sources *[]FlowSource `json:"sources,omitempty"`

	// 状态
	Status *string `json:"status,omitempty"`

	// 流id
	FlowId *string `json:"flow_id,omitempty"`

	// 流输出信息
	Outputs        *[]FlowOutput `json:"outputs,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o ShowFlowDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowFlowDetailResponse struct{}"
	}

	return strings.Join([]string{"ShowFlowDetailResponse", string(data)}, " ")
}
