package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListModifyHistoryResponse Response Object
type ListModifyHistoryResponse struct {

	// 总数。
	TotalCount *int32 `json:"total_count,omitempty"`

	// 参数模板的修改历史列表。
	Histories      *[]ParamGroupHistoryResponse `json:"histories,omitempty"`
	HttpStatusCode int                          `json:"-"`
}

func (o ListModifyHistoryResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListModifyHistoryResponse struct{}"
	}

	return strings.Join([]string{"ListModifyHistoryResponse", string(data)}, " ")
}
