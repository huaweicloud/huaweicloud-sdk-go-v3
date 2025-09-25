package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ImageAppsInfo 查询镜像软件列表，软件信息
type ImageAppsInfo struct {

	// **参数解释**: 软件名称 **取值范围**: 字符长度0-64位
	AppName *string `json:"app_name,omitempty"`

	// **参数解释**: 软件类型 **取值范围**: 字符长度0-64位
	AppType *string `json:"app_type,omitempty"`

	// **参数解释**: 软件版本 **取值范围**: 字符长度0-128位
	AppVersion *string `json:"app_version,omitempty"`

	// **参数解释**: 软件漏洞数 **取值范围**: 最小值0，最大值2147483647
	VulNum *int32 `json:"vul_num,omitempty"`
}

func (o ImageAppsInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImageAppsInfo struct{}"
	}

	return strings.Join([]string{"ImageAppsInfo", string(data)}, " ")
}
