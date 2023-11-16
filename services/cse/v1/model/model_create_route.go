package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateRoute struct {

	// name
	Name *string `json:"name,omitempty"`

	// weight
	Weight *int32 `json:"weight,omitempty"`

	Tags *CreateRouteTags `json:"tags,omitempty"`
}

func (o CreateRoute) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateRoute struct{}"
	}

	return strings.Join([]string{"CreateRoute", string(data)}, " ")
}
