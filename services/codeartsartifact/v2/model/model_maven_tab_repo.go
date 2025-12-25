package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type MavenTabRepo struct {

	// **参数解释**： release仓库名称。  **取值范围**： 不涉及。
	Release *string `json:"release,omitempty"`

	// **参数解释**： snapshot仓库名称。  **取值范围**： 不涉及。
	Snapshot *string `json:"snapshot,omitempty"`
}

func (o MavenTabRepo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MavenTabRepo struct{}"
	}

	return strings.Join([]string{"MavenTabRepo", string(data)}, " ")
}
