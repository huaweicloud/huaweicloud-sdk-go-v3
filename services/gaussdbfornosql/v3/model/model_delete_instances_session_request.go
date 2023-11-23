package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteInstancesSessionRequest Request Object
type DeleteInstancesSessionRequest struct {

	// 节点ID。
	NodeId string `json:"node_id"`

	Body *DeleteInstancesSessionRequestBody `json:"body,omitempty"`
}

func (o DeleteInstancesSessionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteInstancesSessionRequest struct{}"
	}

	return strings.Join([]string{"DeleteInstancesSessionRequest", string(data)}, " ")
}
