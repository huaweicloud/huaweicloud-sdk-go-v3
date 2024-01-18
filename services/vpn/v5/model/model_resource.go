package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Resource struct {
	ResourceId string `json:"resource_id"`

	ResourceDetail *interface{} `json:"resource_detail"`

	Tags []ResourceTag `json:"tags"`

	ResourceName string `json:"resource_name"`
}

func (o Resource) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Resource struct{}"
	}

	return strings.Join([]string{"Resource", string(data)}, " ")
}
