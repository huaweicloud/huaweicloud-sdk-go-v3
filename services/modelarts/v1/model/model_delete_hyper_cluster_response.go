package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteHyperClusterResponse Response Object
type DeleteHyperClusterResponse struct {
	Body           *interface{} `json:"body,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o DeleteHyperClusterResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteHyperClusterResponse struct{}"
	}

	return strings.Join([]string{"DeleteHyperClusterResponse", string(data)}, " ")
}
