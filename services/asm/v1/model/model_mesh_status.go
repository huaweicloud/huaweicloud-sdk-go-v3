package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type MeshStatus struct {

	// 网格状态，取值如下 - Running：运行中，表示网格处于正常运行状态 - Creating：创建中，表示网格正处于创建过程中 - CreateFailed：创建失败 - Deleting：删除中，表示网格正处于删除过程中 - DeleteFailed：删除失败 - Upgrading：升级中，表示网格正处于升级过程中 - UpgradeFailed：升级失败 - RollingBack：回滚中，表示网格正处于回滚过程中 - RollbackFailed：回滚失败
	Phase *string `json:"phase,omitempty"`

	// 网格更新时间
	UpdateTimestamp *sdktime.SdkTime `json:"updateTimestamp,omitempty"`
}

func (o MeshStatus) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MeshStatus struct{}"
	}

	return strings.Join([]string{"MeshStatus", string(data)}, " ")
}
