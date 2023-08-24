package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAppConfigsTemplateResponse Response Object
type ShowAppConfigsTemplateResponse struct {

	// 模板id
	TplId *string `json:"tpl_id,omitempty"`

	// 模板名称
	Name *string `json:"name,omitempty"`

	// 描述
	Description *string `json:"description,omitempty"`

	// 配置项数据
	ConfigTabs *interface{} `json:"config_tabs,omitempty"`

	// 默认数据
	DefaultValues *interface{} `json:"default_values,omitempty"`

	// 创建时间
	CreateTime *string `json:"create_time,omitempty"`

	// 最后一次修改时间
	UpdateTime     *string `json:"update_time,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowAppConfigsTemplateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAppConfigsTemplateResponse struct{}"
	}

	return strings.Join([]string{"ShowAppConfigsTemplateResponse", string(data)}, " ")
}
