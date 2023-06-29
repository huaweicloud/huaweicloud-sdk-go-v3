package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExpandGraph2Request Request Object
type ExpandGraph2Request struct {

	// 图ID。
	GraphId string `json:"graph_id"`

	Body *ExpandGraphReq `json:"body,omitempty"`
}

func (o ExpandGraph2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExpandGraph2Request struct{}"
	}

	return strings.Join([]string{"ExpandGraph2Request", string(data)}, " ")
}
