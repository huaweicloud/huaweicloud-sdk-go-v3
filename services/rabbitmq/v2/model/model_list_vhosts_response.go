package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListVhostsResponse Response Object
type ListVhostsResponse struct {

	// **参数解释**： 当前显示的Vhost数量 **取值范围**： 不涉及。
	Size *int32 `json:"size,omitempty"`

	// **参数解释**： 查询到的Vhost总数 **取值范围**： 不涉及。
	Total *int32 `json:"total,omitempty"`

	// **参数解释**： 查询的Vhost信息列表
	Items          *[]ShowVhostDetailResp `json:"items,omitempty"`
	HttpStatusCode int                    `json:"-"`
}

func (o ListVhostsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListVhostsResponse struct{}"
	}

	return strings.Join([]string{"ListVhostsResponse", string(data)}, " ")
}
