package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// QueryOtTemplateBriefRespDto OT模板概要响应信息
type QueryOtTemplateBriefRespDto struct {

	// 模板id
	TplId *string `json:"tpl_id,omitempty"`

	// 模板名称
	Name *string `json:"name,omitempty"`

	// 描述
	Description *string `json:"description,omitempty"`

	// 创建时间
	CreateTime *string `json:"create_time,omitempty"`

	// 最后一次修改时间
	UpdateTime *string `json:"update_time,omitempty"`
}

func (o QueryOtTemplateBriefRespDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryOtTemplateBriefRespDto struct{}"
	}

	return strings.Join([]string{"QueryOtTemplateBriefRespDto", string(data)}, " ")
}
