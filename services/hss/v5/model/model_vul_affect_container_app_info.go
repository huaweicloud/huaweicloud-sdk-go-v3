package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// VulAffectContainerAppInfo 受漏洞影响容器的软件信息
type VulAffectContainerAppInfo struct {

	// 软件名称
	AppName *string `json:"app_name,omitempty"`

	// 列表
	Paths *[]VulAffectContainerAppPath `json:"paths,omitempty"`

	// 规则
	Rule *string `json:"rule,omitempty"`
}

func (o VulAffectContainerAppInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VulAffectContainerAppInfo struct{}"
	}

	return strings.Join([]string{"VulAffectContainerAppInfo", string(data)}, " ")
}
