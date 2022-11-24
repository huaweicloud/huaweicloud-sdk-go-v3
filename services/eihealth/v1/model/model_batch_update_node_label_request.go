package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type BatchUpdateNodeLabelRequest struct {

	// 节点id
	ServerId string `json:"server_id"`

	Body *BatchUpadteNodeLabelReq `json:"body,omitempty"`
}

func (o BatchUpdateNodeLabelRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUpdateNodeLabelRequest struct{}"
	}

	return strings.Join([]string{"BatchUpdateNodeLabelRequest", string(data)}, " ")
}
