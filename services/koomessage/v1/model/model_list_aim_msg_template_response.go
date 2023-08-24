package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAimMsgTemplateResponse Response Object
type ListAimMsgTemplateResponse struct {

	// 查询结果。
	Result *[]SmsTemplateInfo `json:"result,omitempty"`

	PageInfo       *Page `json:"page_info,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o ListAimMsgTemplateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAimMsgTemplateResponse struct{}"
	}

	return strings.Join([]string{"ListAimMsgTemplateResponse", string(data)}, " ")
}
