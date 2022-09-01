package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type GetTemplatesResponse struct {
	Error *Error `json:"error,omitempty" xml:"error"`

	Result *TemplateRepositoryList `json:"result,omitempty" xml:"result"`

	// 响应状态
	Status         *string `json:"status,omitempty" xml:"status"`
	HttpStatusCode int     `json:"-"`
}

func (o GetTemplatesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetTemplatesResponse struct{}"
	}

	return strings.Join([]string{"GetTemplatesResponse", string(data)}, " ")
}
