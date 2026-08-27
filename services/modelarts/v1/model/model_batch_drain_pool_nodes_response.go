package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDrainPoolNodesResponse Response Object
type BatchDrainPoolNodesResponse struct {
	Body           *interface{} `json:"body,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o BatchDrainPoolNodesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDrainPoolNodesResponse struct{}"
	}

	return strings.Join([]string{"BatchDrainPoolNodesResponse", string(data)}, " ")
}
