package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListAuthorizableTicketsReq struct {

	// 搜索结果
	RequestParams *[]AuthorizableTicketBody `json:"request_params,omitempty"`
}

func (o ListAuthorizableTicketsReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAuthorizableTicketsReq struct{}"
	}

	return strings.Join([]string{"ListAuthorizableTicketsReq", string(data)}, " ")
}
