package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SetStatsConfigBody 设置统计配置请求体
type SetStatsConfigBody struct {

	// **参数解释：** 配置类型 **约束限制：** 不涉及 **取值范围：** - 0：热点统计 - 1：ces上报 **默认取值：** 不涉及
	ConfigType *int32 `json:"config_type,omitempty"`

	// **参数解释：** 资源类型 **约束限制：** 不涉及 **取值范围：** - domain：域名，对应resource_name需配置为域名 - account：账号，对应resource_name需配置为账号 **默认取值：** 不涉及
	ResourceType string `json:"resource_type"`

	// **参数解释：** 资源名称 > 账号或域名  **约束限制：** 不涉及 **取值范围：** 多个资源名称以英文逗号分隔 **默认取值：** 不涉及
	ResourceName string `json:"resource_name"`

	// **参数解释：** 配置信息 **约束限制：** 不涉及 **取值范围：** - ua：HTTP请求头User-Agent的值 - refer：HTTP请求头referer的值 - url：客户访问的http地址 - originurl：回源url **默认取值：** 不涉及
	ConfigInfo *interface{} `json:"config_info"`
}

func (o SetStatsConfigBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetStatsConfigBody struct{}"
	}

	return strings.Join([]string{"SetStatsConfigBody", string(data)}, " ")
}
