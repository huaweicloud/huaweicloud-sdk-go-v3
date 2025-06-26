package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateInstanceRequestBodyResourceTags struct {

	// tag键值.
	Key *string `json:"key,omitempty"`

	// tag值.
	Value *string `json:"value,omitempty"`
}

func (o CreateInstanceRequestBodyResourceTags) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateInstanceRequestBodyResourceTags struct{}"
	}

	return strings.Join([]string{"CreateInstanceRequestBodyResourceTags", string(data)}, " ")
}
