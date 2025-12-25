package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListMavenUserRequest Request Object
type ListMavenUserRequest struct {

	// **参数解释**: 用户名。 **约束限制**: 不涉及。 **取值范围**: 不涉及。 **默认取值**: 不涉及。
	UserName *string `json:"user_name,omitempty"`

	// **参数解释**: 分页查询的页数。 **约束限制**: 不涉及。 **取值范围**: 最小值1，最大值2147483647。 **默认取值**: 1
	PageNo *int32 `json:"page_no,omitempty"`

	// **参数解释**: 分页查询的每页数据量。 **约束限制**: 不涉及。 **取值范围**: 最小值1，最大值100。 **默认取值**: 10
	PageSize *int32 `json:"page_size,omitempty"`
}

func (o ListMavenUserRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListMavenUserRequest struct{}"
	}

	return strings.Join([]string{"ListMavenUserRequest", string(data)}, " ")
}
