package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateImageSyncRepoRequestBody struct {
	// 目标region ID。

	RemoteRegionId string `json:"remoteRegionId"`
	// 目标组织

	RemoteNamespace string `json:"remoteNamespace"`
	// 自动同步，默认为false

	SyncAuto *bool `json:"syncAuto,omitempty"`
	// 是否覆盖，默认为false

	Override *bool `json:"override,omitempty"`
}

func (o CreateImageSyncRepoRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateImageSyncRepoRequestBody struct{}"
	}

	return strings.Join([]string{"CreateImageSyncRepoRequestBody", string(data)}, " ")
}
