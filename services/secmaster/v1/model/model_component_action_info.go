package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ComponentActionInfo 插件执行函数对象
type ComponentActionInfo struct {

	// 插件执行函数id
	Id *string `json:"id,omitempty"`

	// 插件执行函数名称
	ActionName *string `json:"action_name,omitempty"`

	// 插件执行函数描述
	ActionDesc *string `json:"action_desc,omitempty"`

	// 执行函数类型，可选action/connector/manager
	ActionType *string `json:"action_type,omitempty"`

	// 创建时间
	CreateTime *string `json:"create_time,omitempty"`

	// 创建者名称
	CreatorName *string `json:"creator_name,omitempty"`

	// 是否可升级，true/false
	CanUpdate *bool `json:"can_update,omitempty"`

	// 当前启用的插件执行函数版本id
	ActionVersionId *string `json:"action_version_id,omitempty"`

	// 当前启用的用户自定义执行函数版本别名
	ActionVersionName *string `json:"action_version_name,omitempty"`

	// 当前启用的执行函数版本号
	ActionVersionNumber *string `json:"action_version_number,omitempty"`

	// 启用/禁用状态
	ActionEnable *string `json:"action_enable,omitempty"`

	// 当前启用的执行函数版本输入参数列表
	ActionInput *string `json:"action_input,omitempty"`

	// 当前启用的执行函数版本输出参数列表
	ActionOutput *string `json:"action_output,omitempty"`

	// 全量插件执行函数版本列表
	ActionVersions *[]ComponentActionVersionInfo `json:"action_versions,omitempty"`
}

func (o ComponentActionInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ComponentActionInfo struct{}"
	}

	return strings.Join([]string{"ComponentActionInfo", string(data)}, " ")
}
