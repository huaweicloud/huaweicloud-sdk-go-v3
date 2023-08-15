package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteEdgeNodeResponse Response Object
type DeleteEdgeNodeResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteEdgeNodeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteEdgeNodeResponse struct{}"
	}

	return strings.Join([]string{"DeleteEdgeNodeResponse", string(data)}, " ")
}
