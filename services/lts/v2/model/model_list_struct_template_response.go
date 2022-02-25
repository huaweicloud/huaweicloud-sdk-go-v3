package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListStructTemplateResponse struct {
	// 查询的自定义结构化模板数组

	Results        *[]StructTemplateModel `json:"results,omitempty"`
	HttpStatusCode int                    `json:"-"`
}

func (o ListStructTemplateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListStructTemplateResponse struct{}"
	}

	return strings.Join([]string{"ListStructTemplateResponse", string(data)}, " ")
}
