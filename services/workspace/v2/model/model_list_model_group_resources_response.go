package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListModelGroupResourcesResponse Response Object
type ListModelGroupResourcesResponse struct {

	// 总数。
	Total *int32 `json:"total,omitempty"`

	// 应用对象关联详情列表项。
	Items          *[]ModelGroupResourceItemResp `json:"items,omitempty"`
	HttpStatusCode int                           `json:"-"`
}

func (o ListModelGroupResourcesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListModelGroupResourcesResponse struct{}"
	}

	return strings.Join([]string{"ListModelGroupResourcesResponse", string(data)}, " ")
}
