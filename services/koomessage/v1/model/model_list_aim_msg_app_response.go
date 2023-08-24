package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAimMsgAppResponse Response Object
type ListAimMsgAppResponse struct {

	// 查询结果。
	Result *[]SmsApp `json:"result,omitempty"`

	PageInfo       *Page `json:"page_info,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o ListAimMsgAppResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAimMsgAppResponse struct{}"
	}

	return strings.Join([]string{"ListAimMsgAppResponse", string(data)}, " ")
}
