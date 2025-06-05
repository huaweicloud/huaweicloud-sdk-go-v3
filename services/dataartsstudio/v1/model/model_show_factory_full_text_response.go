package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowFactoryFullTextResponse Response Object
type ShowFactoryFullTextResponse struct {

	// 分页大小限制。 取值范围[1,100]。
	Limit *int32 `json:"limit,omitempty"`

	// 当前所在分页。
	Offset *int32 `json:"offset,omitempty"`

	// 查询成功，返回搜索结果。
	SearchDetails *[]SearchDetailV2 `json:"search_details,omitempty"`

	// 成功命中数量。
	TotalHits      *int64 `json:"total_hits,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowFactoryFullTextResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowFactoryFullTextResponse struct{}"
	}

	return strings.Join([]string{"ShowFactoryFullTextResponse", string(data)}, " ")
}
