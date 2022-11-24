package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ResizeGraphRequest struct {

	// 图ID。
	GraphId string `json:"graph_id"`

	Body *ResizeGraphReq `json:"body,omitempty"`
}

func (o ResizeGraphRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResizeGraphRequest struct{}"
	}

	return strings.Join([]string{"ResizeGraphRequest", string(data)}, " ")
}
