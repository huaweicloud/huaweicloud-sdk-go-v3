package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListNamespacesResponse struct {

	// 组织列表
	Namespaces     *[]ShowNamespace `json:"namespaces,omitempty" xml:"namespaces"`
	HttpStatusCode int              `json:"-"`
}

func (o ListNamespacesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListNamespacesResponse struct{}"
	}

	return strings.Join([]string{"ListNamespacesResponse", string(data)}, " ")
}
