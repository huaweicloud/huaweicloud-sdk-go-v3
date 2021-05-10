package model

import (
	"encoding/json"

	"strings"
)

type MrsResource struct {
	// 资源ID

	ResourceId *string `json:"resource_id,omitempty"`
	// 资源详情

	ResourceDetail *string `json:"resource_detail,omitempty"`
	// 标签

	Tags *[]TagPlain `json:"tags,omitempty"`
	// 资源名称

	ResourceName *string `json:"resource_name,omitempty"`
}

func (o MrsResource) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "MrsResource struct{}"
	}

	return strings.Join([]string{"MrsResource", string(data)}, " ")
}
