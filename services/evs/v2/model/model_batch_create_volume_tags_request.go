package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type BatchCreateVolumeTagsRequest struct {
	// 云硬盘ID。

	VolumeId string `json:"volume_id"`

	Body *BatchCreateVolumeTagsRequestBody `json:"body,omitempty"`
}

func (o BatchCreateVolumeTagsRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "BatchCreateVolumeTagsRequest struct{}"
	}

	return strings.Join([]string{"BatchCreateVolumeTagsRequest", string(data)}, " ")
}
