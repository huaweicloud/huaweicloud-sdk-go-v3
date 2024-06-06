package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListNoticeRulesResponse Response Object
type ListNoticeRulesResponse struct {

	// API版本，固定值“v1”，该值不可修改。
	ApiVersion *string `json:"api_version,omitempty"`

	// API类型，固定值“NoticeRule”，该值不可修改。
	Kind *string `json:"kind,omitempty"`

	// 请求数据。
	Items          *[]CreateNoticeRuleRespItem `json:"items,omitempty"`
	HttpStatusCode int                         `json:"-"`
}

func (o ListNoticeRulesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListNoticeRulesResponse struct{}"
	}

	return strings.Join([]string{"ListNoticeRulesResponse", string(data)}, " ")
}
