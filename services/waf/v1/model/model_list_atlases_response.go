package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAtlasesResponse Response Object
type ListAtlasesResponse struct {

	// 总数
	Total *int64 `json:"total,omitempty"`

	// 图库列表
	Items          *[]AtlasItem `json:"items,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ListAtlasesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAtlasesResponse struct{}"
	}

	return strings.Join([]string{"ListAtlasesResponse", string(data)}, " ")
}
