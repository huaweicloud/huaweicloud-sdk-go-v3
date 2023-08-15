package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type OpExtraInfo struct {
	Backup *OpExtendInfoBckup `json:"backup,omitempty"`

	Common *OpExtendInfoCommon `json:"common"`

	Delete *OpExtendInfoDelete `json:"delete,omitempty"`

	Sync *OpExtendInfoSync `json:"sync,omitempty"`

	RemoveResources *OpExtendInfoRemoveResources `json:"remove_resources,omitempty"`

	Replication *OpExtendInfoReplication `json:"replication,omitempty"`

	Resource *Resource `json:"resource"`

	Restore *OpExtendInfoRestore `json:"restore,omitempty"`

	VaultDelete *OpExtendInfoVaultDelete `json:"vault_delete,omitempty"`
}

func (o OpExtraInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OpExtraInfo struct{}"
	}

	return strings.Join([]string{"OpExtraInfo", string(data)}, " ")
}
