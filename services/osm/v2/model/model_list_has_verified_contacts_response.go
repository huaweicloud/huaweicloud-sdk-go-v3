package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListHasVerifiedContactsResponse struct {
	// 总数

	Count *int32 `json:"count,omitempty"`
	// 联系方式列表

	ContactList    *[]ContactV2 `json:"contact_list,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ListHasVerifiedContactsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListHasVerifiedContactsResponse struct{}"
	}

	return strings.Join([]string{"ListHasVerifiedContactsResponse", string(data)}, " ")
}
