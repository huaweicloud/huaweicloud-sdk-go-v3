package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListHasVerifiedContactsResponse struct {

	// 总数
	Count *int32 `json:"count,omitempty" xml:"count"`

	// 联系方式列表
	ContactList    *[]ContactV2 `json:"contact_list,omitempty" xml:"contact_list"`
	HttpStatusCode int          `json:"-"`
}

func (o ListHasVerifiedContactsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListHasVerifiedContactsResponse struct{}"
	}

	return strings.Join([]string{"ListHasVerifiedContactsResponse", string(data)}, " ")
}
