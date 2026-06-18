package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ChangeRequestHookCfgDto struct {

	// **参数解释：** 预留字段，事件触发设置，可为空。
	EventCfgs *[]WebHookEventCfgDto `json:"event_cfgs,omitempty"`

	// **参数解释：** 预留字段，仓库分支规则设置，可为空。
	ProjectCfgs *[]WebHookBranchCfgDto `json:"project_cfgs,omitempty"`

	BranchCfgs *[]WebHookBranchCfgDto `json:"branch_cfgs,omitempty"`
}

func (o ChangeRequestHookCfgDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeRequestHookCfgDto struct{}"
	}

	return strings.Join([]string{"ChangeRequestHookCfgDto", string(data)}, " ")
}
