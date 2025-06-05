package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListUserAppResponse Response Object
type ListUserAppResponse struct {

	// 应用列表。
	Apps *[]AppDto `json:"apps,omitempty"`

	// 应用总条数。
	Count          *int32 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListUserAppResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListUserAppResponse struct{}"
	}

	return strings.Join([]string{"ListUserAppResponse", string(data)}, " ")
}
