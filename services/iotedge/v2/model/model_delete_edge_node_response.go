package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DeleteEdgeNodeResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteEdgeNodeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteEdgeNodeResponse struct{}"
	}

	return strings.Join([]string{"DeleteEdgeNodeResponse", string(data)}, " ")
}
