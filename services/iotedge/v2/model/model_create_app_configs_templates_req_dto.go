package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateAppConfigsTemplatesReqDto 创建应用配置模板请求体
type CreateAppConfigsTemplatesReqDto struct {

	// 模板id
	TplId string `json:"tpl_id"`

	// 模板名称，允许中、数字、英文大小写、下划线、中划线
	Name string `json:"name"`

	// 描述
	Description string `json:"description"`

	// 配置项元数据
	ConfigTabs *interface{} `json:"config_tabs"`

	// 默认配置数据
	DefaultValues *interface{} `json:"default_values,omitempty"`
}

func (o CreateAppConfigsTemplatesReqDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAppConfigsTemplatesReqDto struct{}"
	}

	return strings.Join([]string{"CreateAppConfigsTemplatesReqDto", string(data)}, " ")
}
