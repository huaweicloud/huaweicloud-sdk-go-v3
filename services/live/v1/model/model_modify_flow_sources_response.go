package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ModifyFlowSourcesResponse Response Object
type ModifyFlowSourcesResponse struct {

	// 流id
	FlowId *string `json:"flow_id,omitempty"`

	Source         *FlowSource `json:"source,omitempty"`
	HttpStatusCode int         `json:"-"`
}

func (o ModifyFlowSourcesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyFlowSourcesResponse struct{}"
	}

	return strings.Join([]string{"ModifyFlowSourcesResponse", string(data)}, " ")
}
