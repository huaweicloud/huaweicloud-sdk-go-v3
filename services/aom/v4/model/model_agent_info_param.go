package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AgentInfoParam struct {

	// 分页查询的起始页数（第几页）。默认值：1。
	Page *int64 `json:"page,omitempty"`

	// 每页查询数量，默认20。
	PageSize *int64 `json:"page_size,omitempty"`

	// ecs ID列表信息。
	EcsIdList *[]string `json:"ecs_id_list,omitempty"`

	// agent ID列表信息。
	AgentIdList *[]string `json:"agent_id_list,omitempty"`

	// cmdb  ID列表信息。
	CocCmdbIdList *[]string `json:"coc_cmdb_id_list,omitempty"`
}

func (o AgentInfoParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AgentInfoParam struct{}"
	}

	return strings.Join([]string{"AgentInfoParam", string(data)}, " ")
}
