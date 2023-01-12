package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ImportGraph2Request struct {

	// 图ID。
	GraphId string `json:"graph_id"`

	Body *ImportGraphReq `json:"body,omitempty"`
}

func (o ImportGraph2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImportGraph2Request struct{}"
	}

	return strings.Join([]string{"ImportGraph2Request", string(data)}, " ")
}
