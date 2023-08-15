package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AimPersonalTemplatePage struct {

	// 分页显示，指示当前展示第几页，从1开始，最大支持10页。
	PageOrder int32 `json:"page_order"`

	// 该page下的协议内容。
	Contents []AimPersonalTemplateContent `json:"contents"`
}

func (o AimPersonalTemplatePage) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AimPersonalTemplatePage struct{}"
	}

	return strings.Join([]string{"AimPersonalTemplatePage", string(data)}, " ")
}
