package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListNamespaceTagsRequest Request Object
type ListNamespaceTagsRequest struct {

	// 实例id
	InstanceId string `json:"instance_id"`
}

func (o ListNamespaceTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListNamespaceTagsRequest struct{}"
	}

	return strings.Join([]string{"ListNamespaceTagsRequest", string(data)}, " ")
}
