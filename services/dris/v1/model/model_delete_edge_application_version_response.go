package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteEdgeApplicationVersionResponse Response Object
type DeleteEdgeApplicationVersionResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteEdgeApplicationVersionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteEdgeApplicationVersionResponse struct{}"
	}

	return strings.Join([]string{"DeleteEdgeApplicationVersionResponse", string(data)}, " ")
}
