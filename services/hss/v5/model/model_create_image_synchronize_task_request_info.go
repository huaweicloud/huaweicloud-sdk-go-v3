package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateImageSynchronizeTaskRequestInfo struct {

	// 同步全部仓库镜像 | true 同步全部镜像 false 指定镜像仓同步
	SyncAllRegistries bool `json:"sync_all_registries"`

	// 待同步镜像仓
	RegistryInfo *[]CreateImageSynchronizeTaskRequestInfoRegistryInfo `json:"registry_info,omitempty"`
}

func (o CreateImageSynchronizeTaskRequestInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateImageSynchronizeTaskRequestInfo struct{}"
	}

	return strings.Join([]string{"CreateImageSynchronizeTaskRequestInfo", string(data)}, " ")
}
