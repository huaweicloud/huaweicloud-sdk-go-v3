package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ComponentActionVersionInfo 插件执行函数版本详细信息
type ComponentActionVersionInfo struct {

	// 插件执行函数版本id
	Id *string `json:"id,omitempty"`

	// 创建时间
	CreateTime *string `json:"create_time,omitempty"`

	// 创建者名称
	CreatorName *string `json:"creator_name,omitempty"`

	// 更新时间
	UpdateTime *string `json:"update_time,omitempty"`

	// 用户自定义执行函数版本别名
	ActionVersionName *string `json:"action_version_name,omitempty"`

	// 内部生成的执行函数版本号，只会递增
	ActionVersionNumber *string `json:"action_version_number,omitempty"`

	// 版本描述
	ActionVersionDesc *string `json:"action_version_desc,omitempty"`

	// 执行函数版本输入参数列表
	ActionInput *string `json:"action_input,omitempty"`

	// 执行函数版本输出参数列表
	ActionOutput *string `json:"action_output,omitempty"`

	// 执行函数代码
	ActionCode *string `json:"action_code,omitempty"`

	// 版本审核状态
	ActionStatus *string `json:"action_status,omitempty"`

	// 启用/禁用状态
	ActionEnable *string `json:"action_enable,omitempty"`

	// 调试结果
	DebugResult *string `json:"debug_result,omitempty"`

	// 调试输出
	DebugOutput *string `json:"debug_output,omitempty"`

	// 调试日志
	DebugLog *string `json:"debug_log,omitempty"`
}

func (o ComponentActionVersionInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ComponentActionVersionInfo struct{}"
	}

	return strings.Join([]string{"ComponentActionVersionInfo", string(data)}, " ")
}
