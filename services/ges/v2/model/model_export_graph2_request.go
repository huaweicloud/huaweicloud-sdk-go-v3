package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ExportGraph2Request struct {

	// 图ID。
	GraphId string `json:"graph_id"`

	Body *ExportGraphReq `json:"body,omitempty"`
}

func (o ExportGraph2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportGraph2Request struct{}"
	}

	return strings.Join([]string{"ExportGraph2Request", string(data)}, " ")
}
