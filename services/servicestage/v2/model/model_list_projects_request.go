package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListProjectsRequest struct {

	// 授权名称。
	XRepoAuth string `json:"X-Repo-Auth" xml:"X-Repo-Auth"`

	// 组织ID。
	Namespace string `json:"namespace" xml:"namespace"`
}

func (o ListProjectsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListProjectsRequest struct{}"
	}

	return strings.Join([]string{"ListProjectsRequest", string(data)}, " ")
}
