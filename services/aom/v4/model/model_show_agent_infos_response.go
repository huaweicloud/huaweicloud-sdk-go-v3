package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAgentInfosResponse Response Object
type ShowAgentInfosResponse struct {

	// 页码，默认1
	Page *int64 `json:"page,omitempty"`

	// 每页数量，默认20
	PageSize *int64 `json:"page_size,omitempty"`

	// 总数量
	TotalCount *int64 `json:"total_count,omitempty"`

	// 主机列表信息。
	DataList       *[]AgentInfoResult `json:"data_list,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o ShowAgentInfosResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAgentInfosResponse struct{}"
	}

	return strings.Join([]string{"ShowAgentInfosResponse", string(data)}, " ")
}
