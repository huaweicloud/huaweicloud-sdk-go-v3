package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListVmsTemplateStatusResponseMode 智能信息基础版模板状态查询返回体。
type ListVmsTemplateStatusResponseMode struct {

	// 智能信息基础版模板列表。
	Templates *[]VmsTemplateStatus `json:"templates,omitempty"`

	PageInfo *Page `json:"page_info,omitempty"`
}

func (o ListVmsTemplateStatusResponseMode) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListVmsTemplateStatusResponseMode struct{}"
	}

	return strings.Join([]string{"ListVmsTemplateStatusResponseMode", string(data)}, " ")
}
