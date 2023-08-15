package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTwoTemplatesResponse Response Object
type ListTwoTemplatesResponse struct {
	Error *Error `json:"error,omitempty"`

	Result *TemplateRepositoryList `json:"result,omitempty"`

	// 响应状态
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListTwoTemplatesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTwoTemplatesResponse struct{}"
	}

	return strings.Join([]string{"ListTwoTemplatesResponse", string(data)}, " ")
}
