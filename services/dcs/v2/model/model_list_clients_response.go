package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListClientsResponse Response Object
type ListClientsResponse struct {

	// 数据更新时间
	Time *string `json:"time,omitempty"`

	// 会话列表
	Clients *[]ClientInfo `json:"clients,omitempty"`

	// 会话总数
	Count          *int32 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListClientsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListClientsResponse struct{}"
	}

	return strings.Join([]string{"ListClientsResponse", string(data)}, " ")
}
