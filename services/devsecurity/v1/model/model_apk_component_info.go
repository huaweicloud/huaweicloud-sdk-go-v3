package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ApkComponentInfo struct {

	// 权限列表
	Permission *[]PermissionItem `json:"permission,omitempty"`

	// 安卓activity列表，仅安卓任务存在该组件
	Activity *[]ApkComponentItem `json:"activity,omitempty"`

	// 安卓service列表，仅安卓任务存在该组件
	Service *[]ApkComponentItem `json:"service,omitempty"`

	// 安卓provider列表，仅安卓任务存在该组件
	Provider *[]ApkComponentItem `json:"provider,omitempty"`

	// 安卓receive列表，仅安卓任务存在该组件
	Receive *[]ApkComponentItem `json:"receive,omitempty"`
}

func (o ApkComponentInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApkComponentInfo struct{}"
	}

	return strings.Join([]string{"ApkComponentInfo", string(data)}, " ")
}
