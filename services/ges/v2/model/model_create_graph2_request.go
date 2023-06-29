package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateGraph2Request Request Object
type CreateGraph2Request struct {
	Body *CreateGraphReq `json:"body,omitempty"`
}

func (o CreateGraph2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateGraph2Request struct{}"
	}

	return strings.Join([]string{"CreateGraph2Request", string(data)}, " ")
}
