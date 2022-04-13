package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type EnableDisableEdgeNodesResponse struct {
	Node           *Action `json:"node,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o EnableDisableEdgeNodesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnableDisableEdgeNodesResponse struct{}"
	}

	return strings.Join([]string{"EnableDisableEdgeNodesResponse", string(data)}, " ")
}
