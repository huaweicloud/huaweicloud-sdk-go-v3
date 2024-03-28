package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListJobResourcesRequest Request Object
type ListJobResourcesRequest struct {
	Kind *string `json:"kind,omitempty"`

	Tags *string `json:"tags,omitempty"`
}

func (o ListJobResourcesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListJobResourcesRequest struct{}"
	}

	return strings.Join([]string{"ListJobResourcesRequest", string(data)}, " ")
}
