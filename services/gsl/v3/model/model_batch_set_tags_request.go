package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type BatchSetTagsRequest struct {
	Body *BatchSetTagsReq `json:"body,omitempty" xml:"body"`
}

func (o BatchSetTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchSetTagsRequest struct{}"
	}

	return strings.Join([]string{"BatchSetTagsRequest", string(data)}, " ")
}
