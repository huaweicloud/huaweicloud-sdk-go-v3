package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPublishedTemplatesResponse Response Object
type ListPublishedTemplatesResponse struct {

	// 返回模板的数量。
	Count *int32 `json:"count,omitempty"`

	// 返回模板的列表。
	Templates      *[]TemplateSimpleInfo `json:"templates,omitempty"`
	HttpStatusCode int                   `json:"-"`
}

func (o ListPublishedTemplatesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPublishedTemplatesResponse struct{}"
	}

	return strings.Join([]string{"ListPublishedTemplatesResponse", string(data)}, " ")
}
