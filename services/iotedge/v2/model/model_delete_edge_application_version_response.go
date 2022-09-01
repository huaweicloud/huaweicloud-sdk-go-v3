package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DeleteEdgeApplicationVersionResponse struct {
	Body           *string `json:"body,omitempty" xml:"body"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteEdgeApplicationVersionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteEdgeApplicationVersionResponse struct{}"
	}

	return strings.Join([]string{"DeleteEdgeApplicationVersionResponse", string(data)}, " ")
}
