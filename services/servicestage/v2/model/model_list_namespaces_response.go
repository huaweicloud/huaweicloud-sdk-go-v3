package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListNamespacesResponse Response Object
type ListNamespacesResponse struct {

	// 命名空间列表。
	Namespaces     *[]NamespacesNamespaces `json:"namespaces,omitempty"`
	HttpStatusCode int                     `json:"-"`
}

func (o ListNamespacesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListNamespacesResponse struct{}"
	}

	return strings.Join([]string{"ListNamespacesResponse", string(data)}, " ")
}
