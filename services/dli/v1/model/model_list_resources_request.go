package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListResourcesRequest Request Object
type ListResourcesRequest struct {
	Kind *string `json:"kind,omitempty"`

	Tags *string `json:"tags,omitempty"`
}

func (o ListResourcesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListResourcesRequest struct{}"
	}

	return strings.Join([]string{"ListResourcesRequest", string(data)}, " ")
}
