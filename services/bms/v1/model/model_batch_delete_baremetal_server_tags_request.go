package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteBaremetalServerTagsRequest Request Object
type BatchDeleteBaremetalServerTagsRequest struct {
	ServerId string `json:"server_id"`

	Body *BatchDeleteBaremetalServerTagsRequestBody `json:"body,omitempty"`
}

func (o BatchDeleteBaremetalServerTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteBaremetalServerTagsRequest struct{}"
	}

	return strings.Join([]string{"BatchDeleteBaremetalServerTagsRequest", string(data)}, " ")
}
