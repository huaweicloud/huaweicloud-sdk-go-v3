package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NfsSummary sfsTurbo弹性文件系统输出。
type NfsSummary struct {

	// **参数解释**：sfsTurbo弹性文件系统url。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	NfsPath string `json:"nfs_path"`
}

func (o NfsSummary) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NfsSummary struct{}"
	}

	return strings.Join([]string{"NfsSummary", string(data)}, " ")
}
