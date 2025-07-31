package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListBaselineWhiteListsResponse Response Object
type ListBaselineWhiteListsResponse struct {

	// 基线白名单总数
	TotalNum *int32 `json:"total_num,omitempty"`

	// 基线白名单列表
	DataList *[]BaselineWhiteListsResponseInfo `json:"data_list,omitempty"`

	// 基线检查中检查项的检查类型列表
	TagList        *[]string `json:"tag_list,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListBaselineWhiteListsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBaselineWhiteListsResponse struct{}"
	}

	return strings.Join([]string{"ListBaselineWhiteListsResponse", string(data)}, " ")
}
