package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAccessClientInfosResponse Response Object
type ListAccessClientInfosResponse struct {

	// 接入客户端信息列表
	AccessClientInfos *[]AccessClientInfo `json:"access_client_infos,omitempty"`

	// 接入客户端信息总数
	Total          *int32 `json:"total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListAccessClientInfosResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAccessClientInfosResponse struct{}"
	}

	return strings.Join([]string{"ListAccessClientInfosResponse", string(data)}, " ")
}
