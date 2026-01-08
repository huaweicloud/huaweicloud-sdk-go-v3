package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAddressGroupsDependencyResponse Response Object
type ListAddressGroupsDependencyResponse struct {

	// **参数解释**： 请求ID。 **取值范围**： 不涉及。
	RequestId *string `json:"request_id,omitempty"`

	// **参数解释**： 查询地址组的关联资源的响应体。 **取值范围**： 不涉及。
	AddressGroups *[]AddressGroupDependency `json:"address_groups,omitempty"`

	PageInfo       *PageInfo `json:"page_info,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListAddressGroupsDependencyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAddressGroupsDependencyResponse struct{}"
	}

	return strings.Join([]string{"ListAddressGroupsDependencyResponse", string(data)}, " ")
}
