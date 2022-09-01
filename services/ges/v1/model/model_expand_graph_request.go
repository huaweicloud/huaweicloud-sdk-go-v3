package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ExpandGraphRequest struct {

	// 图ID。
	GraphId string `json:"graph_id" xml:"graph_id"`

	Body *ExpandGraphReq `json:"body,omitempty" xml:"body"`
}

func (o ExpandGraphRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExpandGraphRequest struct{}"
	}

	return strings.Join([]string{"ExpandGraphRequest", string(data)}, " ")
}
