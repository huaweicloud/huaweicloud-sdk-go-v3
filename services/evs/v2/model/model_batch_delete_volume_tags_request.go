package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type BatchDeleteVolumeTagsRequest struct {
	// 磁盘ID。

	VolumeId string `json:"volume_id"`

	Body *BatchDeleteVolumeTagsRequestBody `json:"body,omitempty"`
}

func (o BatchDeleteVolumeTagsRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "BatchDeleteVolumeTagsRequest struct{}"
	}

	return strings.Join([]string{"BatchDeleteVolumeTagsRequest", string(data)}, " ")
}
