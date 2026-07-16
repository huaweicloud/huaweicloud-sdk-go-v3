package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAutoSearchTrialsResponse Response Object
type ShowAutoSearchTrialsResponse struct {

	// 超参搜索所有trial结果的个数。
	Total *int32 `json:"total,omitempty"`

	// 超参搜索所有trial结果的当前页展示个数。
	Count *int32 `json:"count,omitempty"`

	// 超参搜索所有trial结果的当前页展示个数最大值。
	Limit *int32 `json:"limit,omitempty"`

	// 超参搜索所有trial结果的当前页数。
	Offset *int32 `json:"offset,omitempty"`

	// 分类。
	GroupBy *string `json:"group_by,omitempty"`

	Items          *ListAutoSearchTrialsItems `json:"items,omitempty"`
	HttpStatusCode int                        `json:"-"`
}

func (o ShowAutoSearchTrialsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAutoSearchTrialsResponse struct{}"
	}

	return strings.Join([]string{"ShowAutoSearchTrialsResponse", string(data)}, " ")
}
