package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListProtocolsResponse Response Object
type ListProtocolsResponse struct {

	// 请求的唯一标识ID。
	RequestId *string `json:"request_id,omitempty"`

	// 协议列表。
	Protocols      *[]string `json:"protocols,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListProtocolsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListProtocolsResponse struct{}"
	}

	return strings.Join([]string{"ListProtocolsResponse", string(data)}, " ")
}
