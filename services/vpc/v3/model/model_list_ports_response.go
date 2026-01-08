package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPortsResponse Response Object
type ListPortsResponse struct {

	// **参数解释**： 查询端口列表的响应体。 **取值范围**： 不涉及。
	Ports *[]Port `json:"ports,omitempty"`

	// **参数解释**： 请求ID。 **取值范围**： 不涉及。
	RequestId *string `json:"request_id,omitempty"`

	PageInfo       *PageInfo `json:"page_info,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListPortsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPortsResponse struct{}"
	}

	return strings.Join([]string{"ListPortsResponse", string(data)}, " ")
}
