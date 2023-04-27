package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListTemplateDetailsResponse struct {

	// 查询结果
	Results *[]SmsTemplateResp `json:"results,omitempty"`

	// 总数
	Total          *int64 `json:"total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListTemplateDetailsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTemplateDetailsResponse struct{}"
	}

	return strings.Join([]string{"ListTemplateDetailsResponse", string(data)}, " ")
}
