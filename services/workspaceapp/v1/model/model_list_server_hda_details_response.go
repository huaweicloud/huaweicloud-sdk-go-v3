package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListServerHdaDetailsResponse Response Object
type ListServerHdaDetailsResponse struct {

	// 总数。
	Count *int32 `json:"count,omitempty"`

	// 返回列表条目数量上限为分页的最大上限值。
	Items          *[]ServerHdaDetails `json:"items,omitempty"`
	HttpStatusCode int                 `json:"-"`
}

func (o ListServerHdaDetailsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListServerHdaDetailsResponse struct{}"
	}

	return strings.Join([]string{"ListServerHdaDetailsResponse", string(data)}, " ")
}
