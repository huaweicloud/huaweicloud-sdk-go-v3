package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTemplateVersionsResponse Response Object
type ListTemplateVersionsResponse struct {

	// 模板版本列表
	Versions *[]TemplateVersion `json:"versions,omitempty"`

	PageInfo       *PageInfo `json:"page_info,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListTemplateVersionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTemplateVersionsResponse struct{}"
	}

	return strings.Join([]string{"ListTemplateVersionsResponse", string(data)}, " ")
}
