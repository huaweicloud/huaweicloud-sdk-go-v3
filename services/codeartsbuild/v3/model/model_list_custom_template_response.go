package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCustomTemplateResponse Response Object
type ListCustomTemplateResponse struct {
	Result *QueryCustomTemplatesResult `json:"result,omitempty"`

	// 返回错误信息
	Error *string `json:"error,omitempty"`

	// 返回状态信息
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListCustomTemplateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCustomTemplateResponse struct{}"
	}

	return strings.Join([]string{"ListCustomTemplateResponse", string(data)}, " ")
}
