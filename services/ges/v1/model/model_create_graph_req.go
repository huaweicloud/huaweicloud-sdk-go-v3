package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateGraphReq This is a auto create Body Object
type CreateGraphReq struct {
	Graph *Graph `json:"graph"`
}

func (o CreateGraphReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateGraphReq struct{}"
	}

	return strings.Join([]string{"CreateGraphReq", string(data)}, " ")
}
