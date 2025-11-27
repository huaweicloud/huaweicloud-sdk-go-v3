package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListEmailTemplateResponse Response Object
type ListEmailTemplateResponse struct {

	// 总数
	Total *int32 `json:"total,omitempty"`

	// 邮件模板列表
	EmailTemplateList *[]EmailTemplate `json:"email_template_list,omitempty"`
	HttpStatusCode    int              `json:"-"`
}

func (o ListEmailTemplateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEmailTemplateResponse struct{}"
	}

	return strings.Join([]string{"ListEmailTemplateResponse", string(data)}, " ")
}
