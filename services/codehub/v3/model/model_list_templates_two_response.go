package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListTemplatesTwoResponse struct {
	Error *Error `json:"error,omitempty"`
	// 响应结果

	Result *string `json:"result,omitempty"`
	// 响应状态

	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListTemplatesTwoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTemplatesTwoResponse struct{}"
	}

	return strings.Join([]string{"ListTemplatesTwoResponse", string(data)}, " ")
}
