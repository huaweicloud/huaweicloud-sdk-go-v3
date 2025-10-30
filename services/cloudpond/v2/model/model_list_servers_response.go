package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListServersResponse Response Object
type ListServersResponse struct {

	// 服务器列表
	Servers *[]Server `json:"servers,omitempty"`

	PageInfo       *PageInfo `json:"page_info,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListServersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListServersResponse struct{}"
	}

	return strings.Join([]string{"ListServersResponse", string(data)}, " ")
}
