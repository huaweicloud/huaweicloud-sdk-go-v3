package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListRepoAccessoriesResponse Response Object
type ListRepoAccessoriesResponse struct {

	// 附件总数
	Total *int64 `json:"total,omitempty"`

	// 附件列表
	Accessories    *[]AccessoryListAccessories `json:"accessories,omitempty"`
	HttpStatusCode int                         `json:"-"`
}

func (o ListRepoAccessoriesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRepoAccessoriesResponse struct{}"
	}

	return strings.Join([]string{"ListRepoAccessoriesResponse", string(data)}, " ")
}
