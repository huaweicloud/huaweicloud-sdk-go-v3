package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type BatchDeletePublicipTagsRequest struct {
	// 资源ID

	PublicipId string `json:"publicip_id"`

	Body *BatchDeletePublicipTagsRequestBody `json:"body,omitempty"`
}

func (o BatchDeletePublicipTagsRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "BatchDeletePublicipTagsRequest struct{}"
	}

	return strings.Join([]string{"BatchDeletePublicipTagsRequest", string(data)}, " ")
}
