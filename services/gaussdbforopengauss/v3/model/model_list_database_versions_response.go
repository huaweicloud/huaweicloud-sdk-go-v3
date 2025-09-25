package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDatabaseVersionsResponse Response Object
type ListDatabaseVersionsResponse struct {

	// **参数解释**： 数据库版本列表。
	DatabaseVersions *[]DatabaseVersionResult `json:"database_versions,omitempty"`

	// **参数解释**： 数据库三位引擎版本总个数。 **取值范围**： 不涉及。
	Total          *int32 `json:"total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListDatabaseVersionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDatabaseVersionsResponse struct{}"
	}

	return strings.Join([]string{"ListDatabaseVersionsResponse", string(data)}, " ")
}
