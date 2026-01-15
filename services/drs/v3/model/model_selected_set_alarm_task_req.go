package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SelectedSetAlarmTaskReq 需要设置SMN的任务信息。
type SelectedSetAlarmTaskReq struct {

	// 任务ID
	JobId string `json:"job_id"`

	// 任务状态 CREATING:创建中 - CREATE_FAILED:创建失败 - CONFIGURATION:配置中 - STARTJOBING:启动中 - WAITING_FOR_START:等待启动中 - START_JOB_FAILED:启动失败 - PAUSING:已暂停 - FULL_TRANSFER_STARTED:全量开始,灾备场景下为初始化 - FULL_TRANSFER_FAILED:全量失败,灾备场景下为初始化失败 - FULL_TRANSFER_COMPLETE:全量完成,灾备场景下为初始化完成 - INCRE_TRANSFER_STARTED:增量开始,灾备场景下为灾备中 - INCRE_TRANSFER_FAILED:增量失败,灾备场景下为灾备异常 - RELEASE_RESOURCE_STARTED:结束任务中 - RELEASE_RESOURCE_FAILED:结束任务失败 - RELEASE_RESOURCE_COMPLETE:已结束 - REBUILD_NODE_STARTED:故障恢复中 - REBUILD_NODE_FAILED:故障恢复失败 - CHANGE_JOB_STARTED:任务变更中 - CHANGE_JOB_FAILED:任务变更失败 - DELETED:已删除 - CHILD_TRANSFER_STARTING:再编辑子任务启动中 - CHILD_TRANSFER_STARTED:再编辑子任务迁移中 - CHILD_TRANSFER_COMPLETE:再编辑子任务迁移完成 - CHILD_TRANSFER_FAILED:再编辑子任务迁移失败 - RELEASE_CHILD_TRANSFER_STARTED:再编辑子任务结束中 - RELEASE_CHILD_TRANSFER_COMPLETE:再编辑子任务已结束 - NODE_UPGRADE_START:升级开始 - NODE_UPGRADE_COMPLETE:升级完成 - NODE_UPGRADE_FAILED:升级失败
	Status string `json:"status"`

	// 引擎类型
	EngineType string `json:"engine_type"`
}

func (o SelectedSetAlarmTaskReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SelectedSetAlarmTaskReq struct{}"
	}

	return strings.Join([]string{"SelectedSetAlarmTaskReq", string(data)}, " ")
}
