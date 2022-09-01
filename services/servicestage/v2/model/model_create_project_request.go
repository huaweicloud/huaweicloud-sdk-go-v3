package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateProjectRequest struct {

	// 授权名称。
	XRepoAuth string `json:"X-Repo-Auth" xml:"X-Repo-Auth"`

	// 组织ID。
	Namespace string `json:"namespace" xml:"namespace"`

	Body *ProjectCreate `json:"body,omitempty" xml:"body"`
}

func (o CreateProjectRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateProjectRequest struct{}"
	}

	return strings.Join([]string{"CreateProjectRequest", string(data)}, " ")
}
