package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchResizePoolNodesResponse Response Object
type BatchResizePoolNodesResponse struct {
	Body           *interface{} `json:"body,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o BatchResizePoolNodesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchResizePoolNodesResponse struct{}"
	}

	return strings.Join([]string{"BatchResizePoolNodesResponse", string(data)}, " ")
}
