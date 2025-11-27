package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListVendorAccountResponse Response Object
type ListVendorAccountResponse struct {

	// **参数解释：** 云厂商账户列表。 **取值范围：** 不涉及。
	Data *[]BatchListVendorAccountResponseData `json:"data,omitempty"`

	// 总数
	Total *int64 `json:"total,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListVendorAccountResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListVendorAccountResponse struct{}"
	}

	return strings.Join([]string{"ListVendorAccountResponse", string(data)}, " ")
}
