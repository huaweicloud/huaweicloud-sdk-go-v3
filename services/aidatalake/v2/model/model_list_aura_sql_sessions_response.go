package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAuraSqlSessionsResponse Response Object
type ListAuraSqlSessionsResponse struct {

	// **参数解释**：下一页查询marker值，若为空表示当前已是最后一页。 **取值范围**：不涉及。
	NextMarker *string `json:"next_marker,omitempty"`

	// **参数解释**：总数。 **取值范围**：1~2147483647。
	Total *int32 `json:"total,omitempty"`

	// **参数解释**：Session列表。
	Sessions       *[]SqlSessionInfo `json:"sessions,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o ListAuraSqlSessionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAuraSqlSessionsResponse struct{}"
	}

	return strings.Join([]string{"ListAuraSqlSessionsResponse", string(data)}, " ")
}
