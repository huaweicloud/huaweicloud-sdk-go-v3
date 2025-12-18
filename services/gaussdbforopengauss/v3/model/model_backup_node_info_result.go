package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BackupNodeInfoResult struct {

	// **参数解释**: 选择指定az下的节点进行备份。 **取值范围**: 不涉及。
	AzList *string `json:"az_list,omitempty"`

	// **参数解释**: 选择指定节点进行备份。 **取值范围**: 不涉及。
	NodeList *string `json:"node_list,omitempty"`
}

func (o BackupNodeInfoResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BackupNodeInfoResult struct{}"
	}

	return strings.Join([]string{"BackupNodeInfoResult", string(data)}, " ")
}
