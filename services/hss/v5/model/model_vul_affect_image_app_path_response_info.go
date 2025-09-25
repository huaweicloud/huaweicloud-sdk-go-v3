package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// VulAffectImageAppPathResponseInfo 受漏洞影响镜像的软件路径
type VulAffectImageAppPathResponseInfo struct {

	// 处理状态
	Status *string `json:"status,omitempty"`

	// 软件版本
	AppVersion *string `json:"app_version,omitempty"`

	// 软件路径
	AppPath *string `json:"app_path,omitempty"`
}

func (o VulAffectImageAppPathResponseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VulAffectImageAppPathResponseInfo struct{}"
	}

	return strings.Join([]string{"VulAffectImageAppPathResponseInfo", string(data)}, " ")
}
