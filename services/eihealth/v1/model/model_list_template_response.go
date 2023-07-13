package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTemplateResponse Response Object
type ListTemplateResponse struct {

	// 模板总数
	Count *int32 `json:"count,omitempty"`

	// 模板列表
	Templates      *[]TemplateRsp `json:"templates,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ListTemplateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTemplateResponse struct{}"
	}

	return strings.Join([]string{"ListTemplateResponse", string(data)}, " ")
}
