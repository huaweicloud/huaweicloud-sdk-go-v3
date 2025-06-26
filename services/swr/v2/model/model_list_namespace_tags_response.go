package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListNamespaceTagsResponse Response Object
type ListNamespaceTagsResponse struct {

	// 项目下资源标签列表
	Tags           *[]ProjectTag `json:"tags,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o ListNamespaceTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListNamespaceTagsResponse struct{}"
	}

	return strings.Join([]string{"ListNamespaceTagsResponse", string(data)}, " ")
}
