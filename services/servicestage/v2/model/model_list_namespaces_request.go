package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListNamespacesRequest Request Object
type ListNamespacesRequest struct {

	// 授权名称。
	XRepoAuth string `json:"X-Repo-Auth"`
}

func (o ListNamespacesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListNamespacesRequest struct{}"
	}

	return strings.Join([]string{"ListNamespacesRequest", string(data)}, " ")
}
