package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type UpdateEdgeNodeResponse struct {
	Node           *EdgeNodeResp `json:"node,omitempty" xml:"node"`
	HttpStatusCode int           `json:"-"`
}

func (o UpdateEdgeNodeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateEdgeNodeResponse struct{}"
	}

	return strings.Join([]string{"UpdateEdgeNodeResponse", string(data)}, " ")
}
