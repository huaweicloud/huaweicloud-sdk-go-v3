package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Resource struct {
	ResourceId *string `json:"resource_id,omitempty"`

	ResourceDetail *interface{} `json:"resource_detail,omitempty"`

	Tags *[]ResourceTag `json:"tags,omitempty"`

	ResourceName *string `json:"resource_name,omitempty"`
}

func (o Resource) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Resource struct{}"
	}

	return strings.Join([]string{"Resource", string(data)}, " ")
}
