package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PluginControlRequestBody 启动/停止插件请求体
type PluginControlRequestBody struct {

	// **参数解释**： 插件编码 **约束限制**: 不涉及 **取值范围**: 字符长度1-64位 **默认取值**: 不涉及
	PluginCode string `json:"plugin_code"`

	// **参数解释**： 是否是全量操作，优先级高于host_id_list，若取值为true且host_id_list不为空，仍执行全量操作 **约束限制**: 不涉及 **取值范围**: - true：全量操作，启动/停止所有符合条件的插件 - false: 非全量操作，仅对host_id_list内服务器执行启动/停止插件操作 **默认取值**: false
	OperateAll *bool `json:"operate_all,omitempty"`

	// **参数解释**： 服务器ID列表 **约束限制**: 不涉及 **取值范围**: 不涉及 **默认取值**: 不涉及
	HostIdList *[]string `json:"host_id_list,omitempty"`
}

func (o PluginControlRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PluginControlRequestBody struct{}"
	}

	return strings.Join([]string{"PluginControlRequestBody", string(data)}, " ")
}
