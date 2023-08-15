package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResizeGraph2Request Request Object
type ResizeGraph2Request struct {

	// 图ID。
	GraphId string `json:"graph_id"`

	Body *ResizeGraphReq `json:"body,omitempty"`
}

func (o ResizeGraph2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResizeGraph2Request struct{}"
	}

	return strings.Join([]string{"ResizeGraph2Request", string(data)}, " ")
}
