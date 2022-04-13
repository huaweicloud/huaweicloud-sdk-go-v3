package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DeleteEdgeNodeCertsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteEdgeNodeCertsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteEdgeNodeCertsResponse struct{}"
	}

	return strings.Join([]string{"DeleteEdgeNodeCertsResponse", string(data)}, " ")
}
