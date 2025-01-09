package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAppAuthorizationsResponse Response Object
type ListAppAuthorizationsResponse struct {

	// 总数。
	Count *int32 `json:"count,omitempty"`

	// 授权记录。
	Items          *[]AccountInfo `json:"items,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ListAppAuthorizationsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAppAuthorizationsResponse struct{}"
	}

	return strings.Join([]string{"ListAppAuthorizationsResponse", string(data)}, " ")
}
