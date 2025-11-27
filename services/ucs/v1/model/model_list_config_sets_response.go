package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListConfigSetsResponse Response Object
type ListConfigSetsResponse struct {

	// 配置集合列表
	Items *[]ConfigSetResponse `json:"items,omitempty"`

	// 符合查询条件的配置集合总数
	Total          *int32 `json:"total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListConfigSetsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListConfigSetsResponse struct{}"
	}

	return strings.Join([]string{"ListConfigSetsResponse", string(data)}, " ")
}
