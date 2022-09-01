package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type OpExtraInfo struct {
	Backup *OpExtendInfoBckup `json:"backup,omitempty" xml:"backup"`

	Common *OpExtendInfoCommon `json:"common" xml:"common"`

	Delete *OpExtendInfoDelete `json:"delete,omitempty" xml:"delete"`

	Sync *OpExtendInfoSync `json:"sync,omitempty" xml:"sync"`

	RemoveResources *OpExtendInfoRemoveResources `json:"remove_resources,omitempty" xml:"remove_resources"`

	Replication *OpExtendInfoReplication `json:"replication,omitempty" xml:"replication"`

	Resource *Resource `json:"resource" xml:"resource"`

	Restore *OpExtendInfoRestore `json:"restore,omitempty" xml:"restore"`

	VaultDelete *OpExtendInfoVaultDelete `json:"vault_delete,omitempty" xml:"vault_delete"`
}

func (o OpExtraInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OpExtraInfo struct{}"
	}

	return strings.Join([]string{"OpExtraInfo", string(data)}, " ")
}
