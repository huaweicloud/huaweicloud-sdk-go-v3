package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SyncConnectionsNewRequest Request Object
type SyncConnectionsNewRequest struct {
	Body *SyncConnectionsNewRequestBody `json:"body,omitempty"`
}

func (o SyncConnectionsNewRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SyncConnectionsNewRequest struct{}"
	}

	return strings.Join([]string{"SyncConnectionsNewRequest", string(data)}, " ")
}
