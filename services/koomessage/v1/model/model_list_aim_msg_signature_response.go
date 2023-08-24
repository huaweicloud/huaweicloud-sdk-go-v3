package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAimMsgSignatureResponse Response Object
type ListAimMsgSignatureResponse struct {

	// 查询结果。
	Result *[]SmsSignatureInfo `json:"result,omitempty"`

	PageInfo       *Page `json:"page_info,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o ListAimMsgSignatureResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAimMsgSignatureResponse struct{}"
	}

	return strings.Join([]string{"ListAimMsgSignatureResponse", string(data)}, " ")
}
