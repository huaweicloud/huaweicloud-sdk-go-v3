package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EventForensicInfoAutoLaunchForensic **参数解释**： 自启动项取证信息 **取值范围**： 不涉及
type EventForensicInfoAutoLaunchForensic struct {

	// **参数解释**： 事件 **取值范围**： 不涉及
	Event *int32 `json:"event,omitempty"`

	// **参数解释**： 类型 **取值范围**： 不涉及
	Type *int32 `json:"type,omitempty"`

	// **参数解释**： 用户 **取值范围**： 不涉及
	Owner *string `json:"owner,omitempty"`

	// **参数解释**： 命令 **取值范围**： 不涉及
	Name *string `json:"name,omitempty"`

	// **参数解释**： 运行间隔 **取值范围**： 不涉及
	RunInterval *string `json:"run_interval,omitempty"`

	// **参数解释**： hash **取值范围**： 不涉及
	Hash *string `json:"hash,omitempty"`

	// **参数解释**： hash **取值范围**： 不涉及
	Mtime *string `json:"mtime,omitempty"`

	// **参数解释**： 自启动项信息 **取值范围**： 不涉及
	Description *string `json:"description,omitempty"`

	// **参数解释**： 进程文件命令行 **取值范围**： 不涉及
	Path *string `json:"path,omitempty"`
}

func (o EventForensicInfoAutoLaunchForensic) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EventForensicInfoAutoLaunchForensic struct{}"
	}

	return strings.Join([]string{"EventForensicInfoAutoLaunchForensic", string(data)}, " ")
}
