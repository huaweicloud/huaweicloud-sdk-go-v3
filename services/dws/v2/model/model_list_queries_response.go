package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListQueriesResponse Response Object
type ListQueriesResponse struct {

	// **参数解释**： 响应码。 **取值范围**： 不涉及。
	Code *int32 `json:"code,omitempty"`

	// **参数解释**： 响应信息。 **取值范围**： 不涉及。
	Msg *string `json:"msg,omitempty"`

	Data *ListQueriesData `json:"data,omitempty"`

	// **参数解释**： 总条数。 **取值范围**： 不涉及。
	Count          *int32 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListQueriesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListQueriesResponse struct{}"
	}

	return strings.Join([]string{"ListQueriesResponse", string(data)}, " ")
}
