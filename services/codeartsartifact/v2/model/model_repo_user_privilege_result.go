package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RepoUserPrivilegeResult struct {

	// **参数解释**: 用户数量。 **取值范围**: 不涉及。
	Total *int64 `json:"total,omitempty"`

	// **参数解释**: 用户列表。 **取值范围**: 不涉及。
	List *[]RepoUserPrivilegeV5 `json:"list,omitempty"`
}

func (o RepoUserPrivilegeResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RepoUserPrivilegeResult struct{}"
	}

	return strings.Join([]string{"RepoUserPrivilegeResult", string(data)}, " ")
}
