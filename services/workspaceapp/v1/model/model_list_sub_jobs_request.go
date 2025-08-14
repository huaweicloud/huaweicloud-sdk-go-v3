package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSubJobsRequest Request Object
type ListSubJobsRequest struct {

	// job详情的状态： * `WAITING` - 等待 * `RUNNING` - 运行中 * `SUCCESS` - 成功 * `FAILED` - 失败 * `ABNORMAL` - 异常 * `ROLLBACK` - 回滚中 * `ABORTING` - 取消
	Status *string `json:"status,omitempty"`

	// job类型 * `CREATE_SERVER` - 创建服务器 * `DELETE_SERVER` - 删除服务器 * `REJOIN_DOMAIN` - 服务器重新加域 * `CHANGE_SERVER_IMAGE` - 修改服务器镜像 * `REINSTALL_OS` - 服务器重装操作系统 * `MIGRATE_SERVER` - 迁移服务器 * `UPDATE_SERVER_TSVI` - 更新虚拟IP配置 * `UPGRADE_ACCESS_AGENT` - hda升级 * `SCHEDULED_TASK` - 定时任务 * `UPDATE_FREEZE_STATUS` - 更新服务器冻结状态
	JobType string `json:"job_type"`

	// 查询的偏移量，默认值0。
	Offset *int32 `json:"offset,omitempty"`

	// 查询的数量，值区间[1-100]，默认值10。
	Limit *int32 `json:"limit,omitempty"`

	// 任务ID。
	JobId *string `json:"job_id,omitempty"`
}

func (o ListSubJobsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSubJobsRequest struct{}"
	}

	return strings.Join([]string{"ListSubJobsRequest", string(data)}, " ")
}
