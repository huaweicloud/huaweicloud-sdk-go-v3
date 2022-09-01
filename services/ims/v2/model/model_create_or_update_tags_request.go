package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateOrUpdateTagsRequest struct {
	Body *AddOrUpdateTagsRequestBody `json:"body,omitempty" xml:"body"`
}

func (o CreateOrUpdateTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateOrUpdateTagsRequest struct{}"
	}

	return strings.Join([]string{"CreateOrUpdateTagsRequest", string(data)}, " ")
}
