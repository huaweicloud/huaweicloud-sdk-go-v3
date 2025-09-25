package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// WebTamperPrivilegedProcessRequestInfo 特权进程配置详情
type WebTamperPrivilegedProcessRequestInfo struct {

	// **参数解释**: 特权进程文件路径列表 **约束限制** : 不涉及 **取值范围**: 最少0条，最多10条 **默认取值** : 不涉及
	PrivilegedProcessPathList *[]string `json:"privileged_process_path_list,omitempty"`

	// **参数解释**: 特权进程子进程可信状态 **约束限制** : 不涉及 **取值范围**: - True ：开启特权进程子进程可信。 - False ：关闭特权进程子进程可信。  **默认取值** : False
	PrivilegedChildStatus *bool `json:"privileged_child_status,omitempty"`
}

func (o WebTamperPrivilegedProcessRequestInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WebTamperPrivilegedProcessRequestInfo struct{}"
	}

	return strings.Join([]string{"WebTamperPrivilegedProcessRequestInfo", string(data)}, " ")
}
