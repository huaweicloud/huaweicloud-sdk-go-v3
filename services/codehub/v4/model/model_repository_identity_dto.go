package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RepositoryIdentityDto struct {

	// **参数解释：** 项目ID。 **约束限制：** 不涉及。
	Id *int32 `json:"id,omitempty"`

	// **参数解释：** 项目路径。 **约束限制：** 不涉及。
	PathWithNamespace *string `json:"path_with_namespace,omitempty"`
}

func (o RepositoryIdentityDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RepositoryIdentityDto struct{}"
	}

	return strings.Join([]string{"RepositoryIdentityDto", string(data)}, " ")
}
