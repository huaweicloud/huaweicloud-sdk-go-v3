package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// QueryAppConfigsTemplateBriefRespDto 应用配置模板概要响应信息
type QueryAppConfigsTemplateBriefRespDto struct {

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

func (o QueryAppConfigsTemplateBriefRespDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryAppConfigsTemplateBriefRespDto struct{}"
	}

	return strings.Join([]string{"QueryAppConfigsTemplateBriefRespDto", string(data)}, " ")
}
