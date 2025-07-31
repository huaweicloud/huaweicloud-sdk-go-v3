package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AssociateRegistriesResponseInfo struct {

	// 镜像仓库ID
	RegistryId *string `json:"registry_id,omitempty"`

	// 镜像仓库名称
	RegistryName *string `json:"registry_name,omitempty"`

	// **参数解释**: 镜像仓库类型 **取值范围**: - SwrPrivate：swr私有。 - SwrShared：swr共享。 - SwrEnterprise：swr企业。 - Harbor：harbor仓库。 - Jfrog：jfrog仓库。 - Other：其他仓库。
	RegistryType *string `json:"registry_type,omitempty"`

	// 同步状态，包含如下3种。   - finished ：同步完成。   - running ：正在同步。   - failed ：同步失败。
	SyncStatus *string `json:"sync_status,omitempty"`

	// 失败原因
	FailedReason *string `json:"failed_reason,omitempty"`
}

func (o AssociateRegistriesResponseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssociateRegistriesResponseInfo struct{}"
	}

	return strings.Join([]string{"AssociateRegistriesResponseInfo", string(data)}, " ")
}
