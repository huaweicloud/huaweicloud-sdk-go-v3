package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowFlowDetailResponse Response Object
type ShowFlowDetailResponse struct {
	Data           *FlowAnalysisVo `json:"data,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o ShowFlowDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowFlowDetailResponse struct{}"
	}

	return strings.Join([]string{"ShowFlowDetailResponse", string(data)}, " ")
}
