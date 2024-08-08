package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// AppServerTaskStatus server的任务状态: * `scheduling` - 实例处于创建中，正在进行调度 * `block_device_mapping` - 实例处于创建中，正在准备磁盘 * `networking` - 实例处于创建中，正在准备网络 * `spawning` - 实例处于创建中，正在内部创建 * `rebooting` - 实例处于重启中 * `reboot_pending` - 实例处于重启中，正在下发重启。 * `reboot_started` - 实例处于重启中，开始内部重启 * `rebooting_hard` - 实例处于强制重启中 * `reboot_pending_hard` - 实例处于强制重启中，正在下发重启 * `reboot_started_hard` - 实例处于强制重启中，开始内部重启。 * `rebuilding` - 实例处于重建中。 * `rebuild_fail` - 实例重建失败。 * `updating_tsvi` - 实例处于虚拟会话IP更新中。 * `updating_tsvi_failed` - 实例虚拟会话IP更新失败。 * `rebuild_block_device_mapping` - 实例处于重建中，正在准备磁盘。 * `rebuild_spawning` - 实例处于重建中，正在内部重建。 * `migrating` - 实例处于热迁移中。 * `resize_prep` - 实例处于调整规格中，正在准备阶段。 * `resize_migrating` - 实例处于调整规格中，正在迁移阶段。 * `resize_migrated` - 实例处于调整规格中，已经完成迁移。 * `resize_finish` - 实例处于调整规格中，正在完成调整。 * `resize_reverting` - 实例处于调整规格中，正在回退调整。 * `powering-off` - 实例处于停止中。 * `powering-on` - 实例处于启动中。 * `deleting` - 实例处于删除中。 * `source_locking` - 资源锁定中 * `rejoining_domain` - 实例正在重新加域 * `delete_failed` - 实例删除失败 * `upgrading_access_agent` - 实例正在升级AccessAgent * `upgrad_access_agent_fail` - 实例升级AccessAgent失败 * `upgrad_access_agent_success` - 实例升级AccessAgent成功 * `updating_sid` - 实例处于创建中，等待更新SID * `migrate_failed` - 实例迁移失败 * `build_image` - 生成镜像中 * `build_snapshot` - 生成快照中 * `restore_snapshot` - 恢复快照中 * `null` - 未设置
type AppServerTaskStatus struct {
	value string
}

type AppServerTaskStatusEnum struct {
	SCHEDULING                   AppServerTaskStatus
	BLOCK_DEVICE_MAPPING         AppServerTaskStatus
	NETWORKING                   AppServerTaskStatus
	SPAWNING                     AppServerTaskStatus
	REBOOTING                    AppServerTaskStatus
	REBOOT_PENDING               AppServerTaskStatus
	REBOOTING_HARD               AppServerTaskStatus
	REBOOT_PENDING_HARD          AppServerTaskStatus
	REBOOT_STARTED_HARD          AppServerTaskStatus
	REBUILDING                   AppServerTaskStatus
	REBUILD_FAIL                 AppServerTaskStatus
	UPDATING_TSVI                AppServerTaskStatus
	UPDATING_TSVI_FAILED         AppServerTaskStatus
	REBUILD_BLOCK_DEVICE_MAPPING AppServerTaskStatus
	REBUILD_SPAWNING             AppServerTaskStatus
	MIGRATING                    AppServerTaskStatus
	RESIZE_PREP                  AppServerTaskStatus
	RESIZE_MIGRATING             AppServerTaskStatus
	RESIZE_MIGRATED              AppServerTaskStatus
	RESIZE_FINISH                AppServerTaskStatus
	RESIZE_REVERTING             AppServerTaskStatus
	POWERING_OFF                 AppServerTaskStatus
	POWERING_ON                  AppServerTaskStatus
	DELETING                     AppServerTaskStatus
	REJOINING_DOMAIN             AppServerTaskStatus
	SOURCE_LOCKING               AppServerTaskStatus
	DELETE_FAILED                AppServerTaskStatus
	UPDATING_SID                 AppServerTaskStatus
	UPGRADING_ACCESS_AGENT       AppServerTaskStatus
	UPGRAD_ACCESS_AGENT_FAIL     AppServerTaskStatus
	UPGRAD_ACCESS_AGENT_SUCCESS  AppServerTaskStatus
	MIGRATE_FAILED               AppServerTaskStatus
	BUILD_IMAGE                  AppServerTaskStatus
	BUILD_SNAPSHOT               AppServerTaskStatus
	RESTORE_SNAPSHOT             AppServerTaskStatus
	NULL                         AppServerTaskStatus
}

func GetAppServerTaskStatusEnum() AppServerTaskStatusEnum {
	return AppServerTaskStatusEnum{
		SCHEDULING: AppServerTaskStatus{
			value: "scheduling",
		},
		BLOCK_DEVICE_MAPPING: AppServerTaskStatus{
			value: "block_device_mapping",
		},
		NETWORKING: AppServerTaskStatus{
			value: "networking",
		},
		SPAWNING: AppServerTaskStatus{
			value: "spawning",
		},
		REBOOTING: AppServerTaskStatus{
			value: "rebooting",
		},
		REBOOT_PENDING: AppServerTaskStatus{
			value: "reboot_pending",
		},
		REBOOTING_HARD: AppServerTaskStatus{
			value: "rebooting_hard",
		},
		REBOOT_PENDING_HARD: AppServerTaskStatus{
			value: "reboot_pending_hard",
		},
		REBOOT_STARTED_HARD: AppServerTaskStatus{
			value: "reboot_started_hard",
		},
		REBUILDING: AppServerTaskStatus{
			value: "rebuilding",
		},
		REBUILD_FAIL: AppServerTaskStatus{
			value: "rebuild_fail",
		},
		UPDATING_TSVI: AppServerTaskStatus{
			value: "updating_tsvi",
		},
		UPDATING_TSVI_FAILED: AppServerTaskStatus{
			value: "updating_tsvi_failed",
		},
		REBUILD_BLOCK_DEVICE_MAPPING: AppServerTaskStatus{
			value: "rebuild_block_device_mapping",
		},
		REBUILD_SPAWNING: AppServerTaskStatus{
			value: "rebuild_spawning",
		},
		MIGRATING: AppServerTaskStatus{
			value: "migrating",
		},
		RESIZE_PREP: AppServerTaskStatus{
			value: "resize_prep",
		},
		RESIZE_MIGRATING: AppServerTaskStatus{
			value: "resize_migrating",
		},
		RESIZE_MIGRATED: AppServerTaskStatus{
			value: "resize_migrated",
		},
		RESIZE_FINISH: AppServerTaskStatus{
			value: "resize_finish",
		},
		RESIZE_REVERTING: AppServerTaskStatus{
			value: "resize_reverting",
		},
		POWERING_OFF: AppServerTaskStatus{
			value: "powering-off",
		},
		POWERING_ON: AppServerTaskStatus{
			value: "powering-on",
		},
		DELETING: AppServerTaskStatus{
			value: "deleting",
		},
		REJOINING_DOMAIN: AppServerTaskStatus{
			value: "rejoining_domain",
		},
		SOURCE_LOCKING: AppServerTaskStatus{
			value: "source_locking",
		},
		DELETE_FAILED: AppServerTaskStatus{
			value: "delete_failed",
		},
		UPDATING_SID: AppServerTaskStatus{
			value: "updating_sid",
		},
		UPGRADING_ACCESS_AGENT: AppServerTaskStatus{
			value: "upgrading_access_agent",
		},
		UPGRAD_ACCESS_AGENT_FAIL: AppServerTaskStatus{
			value: "upgrad_access_agent_fail",
		},
		UPGRAD_ACCESS_AGENT_SUCCESS: AppServerTaskStatus{
			value: "upgrad_access_agent_success",
		},
		MIGRATE_FAILED: AppServerTaskStatus{
			value: "migrate_failed",
		},
		BUILD_IMAGE: AppServerTaskStatus{
			value: "build_image",
		},
		BUILD_SNAPSHOT: AppServerTaskStatus{
			value: "build_snapshot",
		},
		RESTORE_SNAPSHOT: AppServerTaskStatus{
			value: "restore_snapshot",
		},
		NULL: AppServerTaskStatus{
			value: "null",
		},
	}
}

func (c AppServerTaskStatus) Value() string {
	return c.value
}

func (c AppServerTaskStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AppServerTaskStatus) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}
