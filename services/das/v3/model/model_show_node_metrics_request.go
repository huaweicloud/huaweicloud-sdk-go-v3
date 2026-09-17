package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowNodeMetricsRequest Request Object
type ShowNodeMetricsRequest struct {

	// 节点ID
	NodeId string `json:"node_id"`

	Body *ShowNodeMetricsRequestBody `json:"body,omitempty"`
}

func (o ShowNodeMetricsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowNodeMetricsRequest struct{}"
	}

	return strings.Join([]string{"ShowNodeMetricsRequest", string(data)}, " ")
}
