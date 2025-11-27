package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateClusterResponse Response Object
type UpdateClusterResponse struct {
	Body           *interface{} `json:"body,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o UpdateClusterResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateClusterResponse struct{}"
	}

	return strings.Join([]string{"UpdateClusterResponse", string(data)}, " ")
}
