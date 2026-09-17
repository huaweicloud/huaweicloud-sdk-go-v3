package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListActionsResponse Response Object
type ListActionsResponse struct {

	// **参数解释**： 总条数。 **取值范围**： 大于0。
	Count *int32 `json:"count,omitempty"`

	// **参数解释**： 白名单详情。 **取值范围**： 大于0。
	ActionInfo     *[]ActionDomainInfoDetail `json:"action_info,omitempty"`
	HttpStatusCode int                       `json:"-"`
}

func (o ListActionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListActionsResponse struct{}"
	}

	return strings.Join([]string{"ListActionsResponse", string(data)}, " ")
}
