package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ImportCheckpointRequest Request Object
type ImportCheckpointRequest struct {
	Body *SyncReq `json:"body,omitempty"`
}

func (o ImportCheckpointRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImportCheckpointRequest struct{}"
	}

	return strings.Join([]string{"ImportCheckpointRequest", string(data)}, " ")
}
