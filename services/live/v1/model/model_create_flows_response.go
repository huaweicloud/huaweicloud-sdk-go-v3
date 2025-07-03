package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateFlowsResponse Response Object
type CreateFlowsResponse struct {
	Flow           *FlowDetailRespBody `json:"flow,omitempty"`
	HttpStatusCode int                 `json:"-"`
}

func (o CreateFlowsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateFlowsResponse struct{}"
	}

	return strings.Join([]string{"CreateFlowsResponse", string(data)}, " ")
}
