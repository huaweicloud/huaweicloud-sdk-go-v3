package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListInstanceNamespacesResponse Response Object
type ListInstanceNamespacesResponse struct {

	// 命名空间列表
	Namespaces *[]Namespace `json:"namespaces,omitempty"`

	// 命名空间总数
	Total          *int32 `json:"total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListInstanceNamespacesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstanceNamespacesResponse struct{}"
	}

	return strings.Join([]string{"ListInstanceNamespacesResponse", string(data)}, " ")
}
