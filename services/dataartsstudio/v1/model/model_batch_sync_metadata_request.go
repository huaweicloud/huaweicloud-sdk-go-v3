package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchSyncMetadataRequest Request Object
type BatchSyncMetadataRequest struct {
	Body *CatalogMetaDataEventRequest `json:"body,omitempty"`
}

func (o BatchSyncMetadataRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchSyncMetadataRequest struct{}"
	}

	return strings.Join([]string{"BatchSyncMetadataRequest", string(data)}, " ")
}
