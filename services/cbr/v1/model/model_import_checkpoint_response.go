package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ImportCheckpointResponse Response Object
type ImportCheckpointResponse struct {
	Sync           *SyncRespBody `json:"sync,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o ImportCheckpointResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImportCheckpointResponse struct{}"
	}

	return strings.Join([]string{"ImportCheckpointResponse", string(data)}, " ")
}
