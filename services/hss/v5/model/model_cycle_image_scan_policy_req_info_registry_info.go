package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CycleImageScanPolicyReqInfoRegistryInfo struct {

	// 镜像仓库Id
	RegistryId string `json:"registry_id"`

	// 镜像仓库名称
	RegistryName string `json:"registry_name"`

	// 镜像仓库类型 | SwrPrivate:swr私有 SwrShared:swr共享 SwrEnterprise:swr企业 Harbor:harbor仓库 Jfrog:jfrog仓库 Other:其他仓库
	RegistryType string `json:"registry_type"`
}

func (o CycleImageScanPolicyReqInfoRegistryInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CycleImageScanPolicyReqInfoRegistryInfo struct{}"
	}

	return strings.Join([]string{"CycleImageScanPolicyReqInfoRegistryInfo", string(data)}, " ")
}
