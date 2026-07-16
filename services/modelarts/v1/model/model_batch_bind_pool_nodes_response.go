package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchBindPoolNodesResponse Response Object
type BatchBindPoolNodesResponse struct {
	Body           *interface{} `json:"body,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o BatchBindPoolNodesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchBindPoolNodesResponse struct{}"
	}

	return strings.Join([]string{"BatchBindPoolNodesResponse", string(data)}, " ")
}
