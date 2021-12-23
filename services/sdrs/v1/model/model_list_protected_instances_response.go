package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListProtectedInstancesResponse struct {
	// 保护实例的信息列表。

	ProtectedInstances *[]ShowProtectedInstanceParams `json:"protected_instances,omitempty"`
	// 列表中包含的保护实例个数。

	Count          *int32 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListProtectedInstancesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListProtectedInstancesResponse struct{}"
	}

	return strings.Join([]string{"ListProtectedInstancesResponse", string(data)}, " ")
}
