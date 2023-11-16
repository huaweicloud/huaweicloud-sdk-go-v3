package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListFunctionTemplateResponse Response Object
type ListFunctionTemplateResponse struct {

	// 函数模板列表
	FuncTemplates *[]ShowFunctionTemplateResponseBody `json:"func_templates,omitempty"`

	// 函数下次记录读取位置。
	NextMarker     *int32 `json:"next_marker,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListFunctionTemplateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFunctionTemplateResponse struct{}"
	}

	return strings.Join([]string{"ListFunctionTemplateResponse", string(data)}, " ")
}
