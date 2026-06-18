package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RepositoryTransferBodyDto struct {

	// **参数解释：** 命名空间路径。 **约束限制：** 必填。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	Namespace *string `json:"namespace,omitempty"`
}

func (o RepositoryTransferBodyDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RepositoryTransferBodyDto struct{}"
	}

	return strings.Join([]string{"RepositoryTransferBodyDto", string(data)}, " ")
}
