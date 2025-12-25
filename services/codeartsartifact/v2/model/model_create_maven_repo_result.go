package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateMavenRepoResult struct {

	// **参数解释**： release仓库ID。 **取值范围**： 不涉及。
	Release *string `json:"release,omitempty"`

	// **参数解释**： snapshot仓库ID。 **取值范围**： 不涉及。
	Snapshot *string `json:"snapshot,omitempty"`
}

func (o CreateMavenRepoResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateMavenRepoResult struct{}"
	}

	return strings.Join([]string{"CreateMavenRepoResult", string(data)}, " ")
}
