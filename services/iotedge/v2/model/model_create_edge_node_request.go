package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateEdgeNodeRequest Request Object
type CreateEdgeNodeRequest struct {
	Body *EdgeNodeCreation `json:"body,omitempty"`
}

func (o CreateEdgeNodeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateEdgeNodeRequest struct{}"
	}

	return strings.Join([]string{"CreateEdgeNodeRequest", string(data)}, " ")
}
