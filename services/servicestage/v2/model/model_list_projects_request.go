package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListProjectsRequest Request Object
type ListProjectsRequest struct {

	// 授权名称。
	XRepoAuth string `json:"X-Repo-Auth"`

	// 命名空间ID或者URL编码名称。
	Namespace string `json:"namespace"`
}

func (o ListProjectsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListProjectsRequest struct{}"
	}

	return strings.Join([]string{"ListProjectsRequest", string(data)}, " ")
}
