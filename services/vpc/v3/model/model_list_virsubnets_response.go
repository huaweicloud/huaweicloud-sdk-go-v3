package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListVirsubnetsResponse Response Object
type ListVirsubnetsResponse struct {

	// **参数解释**： 请求ID。 **取值范围**： 不涉及。
	RequestId *string `json:"request_id,omitempty"`

	// **参数解释**： 查询子网列表的响应体。 **取值范围**： 不涉及。
	Virsubnets *[]Virsubnet `json:"virsubnets,omitempty"`

	PageInfo       *PageInfo `json:"page_info,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListVirsubnetsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListVirsubnetsResponse struct{}"
	}

	return strings.Join([]string{"ListVirsubnetsResponse", string(data)}, " ")
}
