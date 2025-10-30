package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// LocalImageAppList repository info
type LocalImageAppList struct {

	// **参数解释**: 软件名称 **取值范围**: 字符长度0-128
	AppName *string `json:"app_name,omitempty"`

	// **参数解释**: 软件类型 **取值范围**: 字符长度0-128
	AppType *string `json:"app_type,omitempty"`

	// **参数解释**: 软件版本 **取值范围**: 字符长度0-128
	AppVersion *string `json:"app_version,omitempty"`

	// **参数解释**: 漏洞总数 **取值范围**: 最小值0，最大值20000
	VulNum *int32 `json:"vul_num,omitempty"`
}

func (o LocalImageAppList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LocalImageAppList struct{}"
	}

	return strings.Join([]string{"LocalImageAppList", string(data)}, " ")
}
