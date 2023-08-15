package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAimSendDetailsResponse Response Object
type ListAimSendDetailsResponse struct {

	// 查询发送明细结果集。
	SendDetails *[]AimSendDetail `json:"send_details,omitempty"`

	PageInfo       *Page `json:"page_info,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o ListAimSendDetailsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAimSendDetailsResponse struct{}"
	}

	return strings.Join([]string{"ListAimSendDetailsResponse", string(data)}, " ")
}
