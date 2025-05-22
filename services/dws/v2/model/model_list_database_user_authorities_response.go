package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDatabaseUserAuthoritiesResponse Response Object
type ListDatabaseUserAuthoritiesResponse struct {

	// **参数解释**： 权限详情列表。 **取值范围**： 不涉及。
	AuthorityList *[]GrantAuthority `json:"authority_list,omitempty"`

	// **参数解释**： 权限总条数。 **取值范围**： 不涉及。
	Count          *int32 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListDatabaseUserAuthoritiesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDatabaseUserAuthoritiesResponse struct{}"
	}

	return strings.Join([]string{"ListDatabaseUserAuthoritiesResponse", string(data)}, " ")
}
