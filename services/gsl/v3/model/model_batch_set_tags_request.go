package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchSetTagsRequest Request Object
type BatchSetTagsRequest struct {
	Body *BatchSetTagsReq `json:"body,omitempty"`
}

func (o BatchSetTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchSetTagsRequest struct{}"
	}

	return strings.Join([]string{"BatchSetTagsRequest", string(data)}, " ")
}
