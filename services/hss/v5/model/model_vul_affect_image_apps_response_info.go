package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// VulAffectImageAppsResponseInfo 受漏洞影响镜像的软件信息
type VulAffectImageAppsResponseInfo struct {

	// 软件名称
	AppName *string `json:"app_name,omitempty"`

	// 列表
	Paths *[]VulAffectImageAppPathResponseInfo `json:"paths,omitempty"`

	// 规则
	Rule *string `json:"rule,omitempty"`
}

func (o VulAffectImageAppsResponseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VulAffectImageAppsResponseInfo struct{}"
	}

	return strings.Join([]string{"VulAffectImageAppsResponseInfo", string(data)}, " ")
}
