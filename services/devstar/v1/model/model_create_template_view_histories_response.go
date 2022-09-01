package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateTemplateViewHistoriesResponse struct {

	// 我浏览的模板。
	Templates *[]TemplateViewHistory `json:"templates,omitempty" xml:"templates"`

	// 我浏览的模板数量。
	Count          *int32 `json:"count,omitempty" xml:"count"`
	HttpStatusCode int    `json:"-"`
}

func (o CreateTemplateViewHistoriesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTemplateViewHistoriesResponse struct{}"
	}

	return strings.Join([]string{"CreateTemplateViewHistoriesResponse", string(data)}, " ")
}
