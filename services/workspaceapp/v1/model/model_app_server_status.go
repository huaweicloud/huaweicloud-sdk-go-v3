package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// AppServerStatus 服务器的稳态，完成某个操作的稳定状态。 * `BUILD` - 创建APS实例,APS实例进入运行之前的状态 * `BUILD_FAIL` - 创建APS实例失败 * `REBOOT` - 实例正在进行重启操作 * `HARD_REBOOT` - 实例正在进行强制重启操作 * `REBUILD` - 实例正在重建中 * `REBUILD_FAIL` - 实例重建失败 * `MIGRATING` - 实例正在热迁移中 * `RESIZE` - 实例接收变更请求，开始进行变更操作 * `ACTIVE` - 实例正常运行状态 * `SHUTOFF` - 实例被正常停止 * `REVERT_RESIZE` - 实例正在回退变更规格的配置 * `VERIFY_RESIZE` - 实例正在校验变更完成后的配置。 * `ERROR` - 实例处于异常状态。 * `DELETING` - 实例删除中。 * `FREEZE` - 冻结 * `BUILD_IMAGE` - 生成镜像中 * `BUILD_SNAPSHOT` - 生成快照中 * `RESTORE_SNAPSHOT` - 恢复快照中 * `NULL` - 未设置
type AppServerStatus struct {
	value string
}

type AppServerStatusEnum struct {
	BUILD            AppServerStatus
	BUILD_FAIL       AppServerStatus
	REBOOT           AppServerStatus
	HARD_REBOOT      AppServerStatus
	REBUILD          AppServerStatus
	REBUILD_FAIL     AppServerStatus
	MIGRATING        AppServerStatus
	RESIZE           AppServerStatus
	ACTIVE           AppServerStatus
	SHUTOFF          AppServerStatus
	REVERT_RESIZE    AppServerStatus
	VERIFY_RESIZE    AppServerStatus
	ERROR            AppServerStatus
	DELETING         AppServerStatus
	FREEZE           AppServerStatus
	BUILD_IMAGE      AppServerStatus
	BUILD_SNAPSHOT   AppServerStatus
	RESTORE_SNAPSHOT AppServerStatus
	NULL             AppServerStatus
}

func GetAppServerStatusEnum() AppServerStatusEnum {
	return AppServerStatusEnum{
		BUILD: AppServerStatus{
			value: "BUILD",
		},
		BUILD_FAIL: AppServerStatus{
			value: "BUILD_FAIL",
		},
		REBOOT: AppServerStatus{
			value: "REBOOT",
		},
		HARD_REBOOT: AppServerStatus{
			value: "HARD_REBOOT",
		},
		REBUILD: AppServerStatus{
			value: "REBUILD",
		},
		REBUILD_FAIL: AppServerStatus{
			value: "REBUILD_FAIL",
		},
		MIGRATING: AppServerStatus{
			value: "MIGRATING",
		},
		RESIZE: AppServerStatus{
			value: "RESIZE",
		},
		ACTIVE: AppServerStatus{
			value: "ACTIVE",
		},
		SHUTOFF: AppServerStatus{
			value: "SHUTOFF",
		},
		REVERT_RESIZE: AppServerStatus{
			value: "REVERT_RESIZE",
		},
		VERIFY_RESIZE: AppServerStatus{
			value: "VERIFY_RESIZE",
		},
		ERROR: AppServerStatus{
			value: "ERROR",
		},
		DELETING: AppServerStatus{
			value: "DELETING",
		},
		FREEZE: AppServerStatus{
			value: "FREEZE",
		},
		BUILD_IMAGE: AppServerStatus{
			value: "BUILD_IMAGE",
		},
		BUILD_SNAPSHOT: AppServerStatus{
			value: "BUILD_SNAPSHOT",
		},
		RESTORE_SNAPSHOT: AppServerStatus{
			value: "RESTORE_SNAPSHOT",
		},
		NULL: AppServerStatus{
			value: "null",
		},
	}
}

func (c AppServerStatus) Value() string {
	return c.value
}

func (c AppServerStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AppServerStatus) UnmarshalJSON(b []byte) error {
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
