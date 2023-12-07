package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// OfflineNodesResponse Response Object
type OfflineNodesResponse struct {
	Body           *interface{} `json:"body,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o OfflineNodesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OfflineNodesResponse struct{}"
	}

	return strings.Join([]string{"OfflineNodesResponse", string(data)}, " ")
}
