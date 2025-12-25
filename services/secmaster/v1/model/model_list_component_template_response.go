package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListComponentTemplateResponse Response Object
type ListComponentTemplateResponse struct {

	// 组件模板数量
	Count *int64 `json:"count,omitempty"`

	// 组件模板
	Records        *[]ListComponentTemplate `json:"records,omitempty"`
	HttpStatusCode int                      `json:"-"`
}

func (o ListComponentTemplateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListComponentTemplateResponse struct{}"
	}

	return strings.Join([]string{"ListComponentTemplateResponse", string(data)}, " ")
}
