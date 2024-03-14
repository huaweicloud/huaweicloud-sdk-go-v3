package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListReceivedHandshakesResponse Response Object
type ListReceivedHandshakesResponse struct {

	// 邀请（握手）对象的列表，其中包含与指定账号关联的每个邀请（握手）的详细信息。
	Handshakes *[]HandshakeDto `json:"handshakes,omitempty"`

	PageInfo       *PageInfoDto `json:"page_info,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ListReceivedHandshakesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListReceivedHandshakesResponse struct{}"
	}

	return strings.Join([]string{"ListReceivedHandshakesResponse", string(data)}, " ")
}
