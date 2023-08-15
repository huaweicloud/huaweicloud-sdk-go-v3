package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListConformancePacksResponse Response Object
type ListConformancePacksResponse struct {

	// 合规规则包查询列表。
	Value *[]ConformancePack `json:"value,omitempty"`

	PageInfo       *PageInfo `json:"page_info,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListConformancePacksResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListConformancePacksResponse struct{}"
	}

	return strings.Join([]string{"ListConformancePacksResponse", string(data)}, " ")
}
