package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowTakeOverAssetDetailsRequest struct {
	// 媒资原始输入存放的桶<br/>

	SourceBucket string `json:"source_bucket"`
	// 媒资原始输入的objectKey。<br/>

	SourceObject string `json:"source_object"`
}

func (o ShowTakeOverAssetDetailsRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowTakeOverAssetDetailsRequest struct{}"
	}

	return strings.Join([]string{"ShowTakeOverAssetDetailsRequest", string(data)}, " ")
}
