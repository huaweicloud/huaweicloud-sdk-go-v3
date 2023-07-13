package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateEdgeNodeResponse Response Object
type CreateEdgeNodeResponse struct {
	Node           *EdgeNodeResp `json:"node,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o CreateEdgeNodeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateEdgeNodeResponse struct{}"
	}

	return strings.Join([]string{"CreateEdgeNodeResponse", string(data)}, " ")
}
