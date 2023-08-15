package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateProjectRequest Request Object
type CreateProjectRequest struct {

	// 授权名称。
	XRepoAuth string `json:"X-Repo-Auth"`

	// 命名空间ID或者URL编码名称。
	Namespace string `json:"namespace"`

	Body *ProjectCreate `json:"body,omitempty"`
}

func (o CreateProjectRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateProjectRequest struct{}"
	}

	return strings.Join([]string{"CreateProjectRequest", string(data)}, " ")
}
