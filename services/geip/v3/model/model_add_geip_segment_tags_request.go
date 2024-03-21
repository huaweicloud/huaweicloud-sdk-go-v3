package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddGeipSegmentTagsRequest Request Object
type AddGeipSegmentTagsRequest struct {
	ResourceId string `json:"resource_id"`

	Body *CreateV2TagRequestBody `json:"body,omitempty"`
}

func (o AddGeipSegmentTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddGeipSegmentTagsRequest struct{}"
	}

	return strings.Join([]string{"AddGeipSegmentTagsRequest", string(data)}, " ")
}
