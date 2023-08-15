package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTemplateVarilableDetailsResponse Response Object
type ListTemplateVarilableDetailsResponse struct {

	// 查询结果
	Results *[]ApiTemplateVariable `json:"results,omitempty"`

	// 总数
	Total          *int64 `json:"total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListTemplateVarilableDetailsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTemplateVarilableDetailsResponse struct{}"
	}

	return strings.Join([]string{"ListTemplateVarilableDetailsResponse", string(data)}, " ")
}
