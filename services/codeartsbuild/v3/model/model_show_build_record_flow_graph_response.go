package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowBuildRecordFlowGraphResponse Response Object
type ShowBuildRecordFlowGraphResponse struct {

	// 状态
	Status *string `json:"status,omitempty"`

	// 错误
	Error *interface{} `json:"error,omitempty"`

	Result         *FlowGraph2Result `json:"result,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o ShowBuildRecordFlowGraphResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowBuildRecordFlowGraphResponse struct{}"
	}

	return strings.Join([]string{"ShowBuildRecordFlowGraphResponse", string(data)}, " ")
}
