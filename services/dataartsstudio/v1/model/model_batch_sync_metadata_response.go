package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchSyncMetadataResponse Response Object
type BatchSyncMetadataResponse struct {

	// 是否成功
	Success        *bool `json:"success,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o BatchSyncMetadataResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchSyncMetadataResponse struct{}"
	}

	return strings.Join([]string{"BatchSyncMetadataResponse", string(data)}, " ")
}
