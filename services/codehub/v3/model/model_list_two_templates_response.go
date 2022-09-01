package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListTwoTemplatesResponse struct {
	Error *Error `json:"error,omitempty" xml:"error"`

	Result *TemplateRepositoryList `json:"result,omitempty" xml:"result"`

	// 响应状态
	Status         *string `json:"status,omitempty" xml:"status"`
	HttpStatusCode int     `json:"-"`
}

func (o ListTwoTemplatesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTwoTemplatesResponse struct{}"
	}

	return strings.Join([]string{"ListTwoTemplatesResponse", string(data)}, " ")
}
