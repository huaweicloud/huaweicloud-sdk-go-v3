package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SnapshotSpecItems struct {

	// 子任务ID
	Id *string `json:"id,omitempty"`

	// **参数解释**： 子任务类型 **取值范围**： - master-backup：集群EVS备份 - master-backup-rollback：集群EVS回滚 **默认取值**： 不涉及
	Type *string `json:"type,omitempty"`

	// **参数解释**： 状态 **取值范围**： - Init：初始化 - Queuing：等待 - Running：运行中 - Pause：暂停 - Success：成功 - Failed：失败 **默认取值**： 不涉及
	Status *string `json:"status,omitempty"`

	// 任务创建时间
	CreationTimestamp *string `json:"creationTimestamp,omitempty"`

	// 任务更新时间
	UpdateTimestamp *string `json:"updateTimestamp,omitempty"`

	// 信息
	Message *string `json:"message,omitempty"`
}

func (o SnapshotSpecItems) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SnapshotSpecItems struct{}"
	}

	return strings.Join([]string{"SnapshotSpecItems", string(data)}, " ")
}
