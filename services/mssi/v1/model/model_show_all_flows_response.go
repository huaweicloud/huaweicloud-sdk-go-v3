package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAllFlowsResponse Response Object
type ShowAllFlowsResponse struct {

	// 流的数量
	Count *string `json:"count,omitempty"`

	// 流列表
	FlowMetas      *[]FlowMeta `json:"flow_metas,omitempty"`
	HttpStatusCode int         `json:"-"`
}

func (o ShowAllFlowsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAllFlowsResponse struct{}"
	}

	return strings.Join([]string{"ShowAllFlowsResponse", string(data)}, " ")
}
