package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// WorkitemStatusStatus 工作项的状态
type WorkitemStatusStatus struct {

	// 工作项的状态id
	Id *string `json:"id,omitempty"`

	// 状态名称
	Name *string `json:"name,omitempty"`

	// 工作项状态的类型， BACKLOG( \"初始化\"), READY(\"待启动\"), IN_PROGRESS(\"进行中\"), COMPLETE(\"已完成\"), DONE(\"已结束\"),
	Type *string `json:"type,omitempty"`

	// 工作项状态的描述
	Description *string `json:"description,omitempty"`

	// 工作项状态的父状态id
	ParentStatusId *string `json:"parent_status_id,omitempty"`
}

func (o WorkitemStatusStatus) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WorkitemStatusStatus struct{}"
	}

	return strings.Join([]string{"WorkitemStatusStatus", string(data)}, " ")
}
