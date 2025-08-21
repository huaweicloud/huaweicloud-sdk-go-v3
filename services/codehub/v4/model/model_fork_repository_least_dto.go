package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ForkRepositoryLeastDto struct {

	// **参数解释：** 带命名空间的仓库路径。 **约束限制：** view=least时返回
	PathWithNamespace *string `json:"path_with_namespace,omitempty"`
}

func (o ForkRepositoryLeastDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ForkRepositoryLeastDto struct{}"
	}

	return strings.Join([]string{"ForkRepositoryLeastDto", string(data)}, " ")
}
