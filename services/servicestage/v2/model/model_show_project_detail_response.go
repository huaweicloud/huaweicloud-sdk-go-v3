package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowProjectDetailResponse struct {

	// 命名空间ID。
	NamespaceId *string `json:"namespace_id,omitempty" xml:"namespace_id"`

	// 命名空间。
	Namespace *string `json:"namespace,omitempty" xml:"namespace"`

	// 仓库项目ID。
	ProjectId *string `json:"project_id,omitempty" xml:"project_id"`

	// 仓库项目。
	Project        *string `json:"project,omitempty" xml:"project"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowProjectDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowProjectDetailResponse struct{}"
	}

	return strings.Join([]string{"ShowProjectDetailResponse", string(data)}, " ")
}
