package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteEdgeAppResponse Response Object
type DeleteEdgeAppResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteEdgeAppResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteEdgeAppResponse struct{}"
	}

	return strings.Join([]string{"DeleteEdgeAppResponse", string(data)}, " ")
}
