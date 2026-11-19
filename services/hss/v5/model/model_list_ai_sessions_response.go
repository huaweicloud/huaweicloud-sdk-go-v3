package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAiSessionsResponse Response Object
type ListAiSessionsResponse struct {

	// **参数解释**： AI会话总数 **取值范围**： 0~2000000
	TotalNum *int32 `json:"total_num,omitempty"`

	// **参数解释**： AI会话列表 **取值范围**： 不涉及
	DataList       *[]ListAiSessionsResponseBodyData `json:"data_list,omitempty"`
	HttpStatusCode int                               `json:"-"`
}

func (o ListAiSessionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAiSessionsResponse struct{}"
	}

	return strings.Join([]string{"ListAiSessionsResponse", string(data)}, " ")
}
