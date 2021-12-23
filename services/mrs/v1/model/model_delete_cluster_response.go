package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DeleteClusterResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteClusterResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteClusterResponse struct{}"
	}

	return strings.Join([]string{"DeleteClusterResponse", string(data)}, " ")
}
