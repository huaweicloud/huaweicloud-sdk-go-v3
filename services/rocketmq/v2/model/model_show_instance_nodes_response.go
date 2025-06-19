package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowInstanceNodesResponse Response Object
type ShowInstanceNodesResponse struct {

	// 下个分页的offset。
	NextOffset *int32 `json:"next_offset,omitempty"`

	// 上个分页的offset。
	PreviousOffset *int32 `json:"previous_offset,omitempty"`

	// 后台任务ID
	Nodes *[]NodeContextEntity `json:"nodes,omitempty"`

	// 总个数
	Total          *int32 `json:"total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowInstanceNodesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowInstanceNodesResponse struct{}"
	}

	return strings.Join([]string{"ShowInstanceNodesResponse", string(data)}, " ")
}
