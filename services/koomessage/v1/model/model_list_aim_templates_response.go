package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAimTemplatesResponse Response Object
type ListAimTemplatesResponse struct {

	// 模板列表。
	Templates *[]AimTemplate `json:"templates,omitempty"`

	PageInfo       *Page `json:"page_info,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o ListAimTemplatesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAimTemplatesResponse struct{}"
	}

	return strings.Join([]string{"ListAimTemplatesResponse", string(data)}, " ")
}
